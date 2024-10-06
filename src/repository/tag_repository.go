package repository

import (
	"momonga_blog/database"
	"momonga_blog/repository/model"
)


type TagRepositoryInterface interface {
	GetTags() ([]*model.Tag, error)
	GetTagByBlogUuid(blogId uint) ([]*model.Tag, error)
	GetTagsByUuids(tagUuids []string) ([]*model.Tag, error)
}

type TagRepository struct {}

var _ TagRepositoryInterface = &TagRepository{}

func NewTagRepository() *TagRepository {
	return &TagRepository{}
}

func (tr *TagRepository) GetTags() ([]*model.Tag, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	var tags []*model.Tag
	result := db.Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}

	return tags, nil
}

func (tr *TagRepository) GetTagByBlogUuid(blogId uint) ([]*model.Tag, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	var tags []*model.Tag
	result := db.Model(&model.Tag{}).
        Joins("inner join blog_tags on blog_tags.tag_id = tags.id").
        Where("blog_tags.blog_id = ?", blogId).
        Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}

	return tags, nil
}

func (tr *TagRepository) GetTagsByUuids(tagUuids []string) ([]*model.Tag, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	var tags []*model.Tag
	result := db.Where("uuid IN (?)", tagUuids).Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}

	return tags, nil
}