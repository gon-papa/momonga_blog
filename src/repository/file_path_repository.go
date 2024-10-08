package repository

import (
	"momonga_blog/database"
	"momonga_blog/repository/model"
)

type FilePathRepositoryInterface interface {
	CreateFilePath(filePath string, name string) (*model.FilePath, error)
}

type FilePathRepository struct {}

var _ FilePathRepositoryInterface = &FilePathRepository{}

func NewFilePathRepository() *FilePathRepository {
	return &FilePathRepository{}
}

func (fpr *FilePathRepository) CreateFilePath(filePath string, name string) (*model.FilePath, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}
	
	filePathObj := &model.FilePath{
		Name: name,
		FilePath: filePath,
	}

	result := db.Create(filePathObj)
	if result.Error != nil {
		return nil, result.Error
	}

	return filePathObj, nil
}