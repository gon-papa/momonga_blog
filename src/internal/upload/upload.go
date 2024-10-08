package upload

import (
	"io"
	"momonga_blog/config"
)

type Uploader interface {
	// GetEndpoint() string
	UploadImage(file io.Reader, name string) (string, error)
}


func UploaderSelector() Uploader {
	cnf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	if cnf.Env == "production" {
		// TODO: implement S3Upload
		return &LocalUpload{}
	}
	return &LocalUpload{}
}