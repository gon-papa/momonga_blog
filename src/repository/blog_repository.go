package repository

import (
	"momonga_blog/database"
	"momonga_blog/internal/types"
	"momonga_blog/repository/model"
)


type BlogRepositoryInterface interface {
	GetBlogs(page types.Page, limit types.Limit) (*types.BlogList, error)
	FindBlogByUUID(uuid string) (*model.Blog, error)
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

func (br *BlogRepository) FindBlogByUUID(uuid string) (*model.Blog, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	result := db.Where("uuid = ?", uuid).First(&br.model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &br.model, nil
}