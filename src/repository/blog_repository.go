package repository

import (
	"momonga_blog/database"
	"momonga_blog/internal/logging"
	"momonga_blog/internal/types"
	"momonga_blog/repository/model"
	"time"

	"github.com/google/uuid"
)


type BlogRepositoryInterface interface {
	GetBlogs(page types.Page, limit types.Limit) (*types.BlogList, error)
	FindBlogByUUID(uuid types.Uuid) (*model.Blog, error)
	CreateBlog(blog types.CreateBlogData, tagUuids []*model.Tag) (*model.Blog, error) 
	UpdateBlog(uuid types.Uuid, blog types.UpdateBlogData, tagUuids []string) (*model.Blog, error)
}

type BlogRepository struct {}

var _ BlogRepositoryInterface = &BlogRepository{}

func NewBlogRepository() *BlogRepository {
	return &BlogRepository{}
}

func (br *BlogRepository) GetBlogs(page types.Page, limit types.Limit) (*types.BlogList, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	// 全体のレコード数を取得
	var total int64
	countResult := db.Model(&model.Blog{}).Count(&total)
	if countResult.Error != nil {
		return nil, countResult.Error
	}

	var blogs []*model.Blog

	result := db.Preload("Tags").Offset((page.ToInt() - 1) * limit.ToInt()).Limit(limit.ToInt()).Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}

	blogList := types.NewBlogList(blogs, page.ToInt(), limit.ToInt(), total)

	return blogList, nil
}

func (br *BlogRepository) FindBlogByUUID(uuid types.Uuid) (*model.Blog, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	var blog model.Blog
	result := db.Preload("Tags").Where("uuid = ?", uuid.ToString()).First(&blog)
	if result.Error != nil {
		return nil, result.Error
	}
	return &blog, nil
}

func (br *BlogRepository) CreateBlog(blog types.CreateBlogData, tags []*model.Tag) (*model.Blog, error)  {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	var saveBlog *model.Blog = &model.Blog{
		UUID: uuid.New().String(),
		Year: int(time.Now().Year()),
		Month: int(time.Now().Month()),
		Day: int(time.Now().Day()),
		Title: blog.Title,
		Body: blog.Body,
		IsShow: blog.IsShow,
		DeletedAt: nil,
		Tags: tags,
	}

	//  トランザクション
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Create(saveBlog)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	return saveBlog, nil
}


func (br *BlogRepository) UpdateBlog(uuid types.Uuid, blog types.UpdateBlogData, tags []string) (*model.Blog, error) {
    db, err := database.GetDB()
    if err != nil {
        return nil, err
    }

    // トランザクション開始
    tx := db.Begin()
    if tx.Error != nil {
        return nil, tx.Error
    }
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    var updateBlog model.Blog
    result := tx.Where("uuid = ?", uuid.ToString()).First(&updateBlog)
    if result.Error != nil {
        tx.Rollback()
        return nil, result.Error
    }

    // 新しいタグを取得
    var newTags []*model.Tag
    if len(tags) > 0 {
        if err := tx.Where("uuid IN ?", tags).Find(&newTags).Error; err != nil {
            tx.Rollback()
            logging.ErrorLogger.Error("Failed to get tags", "error", err)
            return nil, err
        }
    } else {
        newTags = []*model.Tag{}
    }

	// タグの関連付けをクリア
	if err := tx.Model(&updateBlog).Association("Tags").Clear(); err != nil {
		tx.Rollback()
		return nil, err
	}

	// ブログの情報を更新
	updateBlog.Title = blog.Title
	updateBlog.Body = blog.Body
	updateBlog.IsShow = blog.IsShow
	updateBlog.Tags = newTags

    // ブログを保存
    if err := tx.Save(&updateBlog).Error; err != nil {
        tx.Rollback()
        return nil, err
    }

    // トランザクションをコミット
    if err := tx.Commit().Error; err != nil {
        return nil, err
    }

    return &updateBlog, nil
}