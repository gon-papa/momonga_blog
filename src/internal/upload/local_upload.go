package upload

import (
	"io"
	"momonga_blog/config"
	"momonga_blog/consts"
	"os"
)

type LocalUpload struct {}

var _ Uploader = &LocalUpload{}

func NewLocalUpload() Uploader {
	return &LocalUpload{}
}

func (l *LocalUpload) UploadImage(file io.Reader, name string) (string, error) {
	savePath, err := l.saveUploadFile(file, name)
	if err != nil {
		return "", err
	}
	
	return savePath, nil
}

func (l *LocalUpload) saveUploadFile(file io.Reader, name string) (string, error) {
	dir := consts.ImageSaveDir
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}
	}
	
	savePath := dir + name
	out, err := os.Create(savePath)
	if err != nil {
		return "", err
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	cnf, err := config.GetConfig()
	if err != nil {
		return "", err
	}
	getPath := cnf.Url + cnf.Port + consts.ImageFileEndpoint + name

	return getPath, nil
}