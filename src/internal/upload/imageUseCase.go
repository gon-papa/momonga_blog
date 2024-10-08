package upload

import (
	"io"
	"momonga_blog/repository"
	"momonga_blog/repository/model"
)

type ImageUseCase interface {
	UploadImage(file io.Reader, name string) (*model.FilePath, error)
}

type imageUseCase struct {
	uploader Uploader
	repository repository.FilePathRepositoryInterface
}

var _ ImageUseCase = &imageUseCase{}

func NewImageUseCase() ImageUseCase {
	return &imageUseCase{
		uploader: UploaderSelector(),
		repository: repository.NewFilePathRepository(),
	}
}

func (iuc *imageUseCase) UploadImage(file io.Reader, name string) (*model.FilePath, error) {
	// ファイルを保存
	filePath, err := iuc.uploader.UploadImage(file, name)
	if err != nil {
		return nil, err
	}

	// ファイルのパスをDBに保存
	filePathObj, err := iuc.repository.CreateFilePath(filePath, name)
	if err != nil {
		return nil, err
	}

	return filePathObj, nil
}