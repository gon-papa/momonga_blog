package repository

import (
	"momonga_blog/database"
	"momonga_blog/repository/model"
)


type TagRepositoryInterface interface {
	GetTagsByUuids(tagUuids []string) ([]*model.Tag, error)
}

type TagRepository struct {}

var _ TagRepositoryInterface = &TagRepository{}

func NewTagRepository() *TagRepository {
	return &TagRepository{}
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