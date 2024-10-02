package repository

import (
	"momonga_blog/database"
	"momonga_blog/internal/types"
	"momonga_blog/repository/model"
	"time"

	"github.com/google/uuid"
)


type BlogRepositoryInterface interface {
	GetBlogs(page types.Page, limit types.Limit) (*types.BlogList, error)
	FindBlogByUUID(uuid types.Uuid) (*model.Blog, error)
	CreateBlog(blog types.CreateBlogData, tags [] types.CreateTagData) (*model.Blog, error) 
}

type BlogRepository struct {
	model model.Blog
}

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

	result := db.Where("uuid = ?", uuid.ToString()).First(&br.model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &br.model, nil
}

func (br *BlogRepository) CreateBlog(blog types.CreateBlogData, tags []types.CreateTagData) (*model.Blog, error)  {
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
	 }

	 if len(tags) > 0 {
		for _, tag := range tags {
			saveBlog.Tags = append(saveBlog.Tags, model.Tag{
				UUID: uuid.New().String(),
				Name: tag.Name,
			})
		}
	}
	 
	result := db.Create(saveBlog)
	if result.Error != nil {
		return nil, result.Error
	}

	return saveBlog, nil
}