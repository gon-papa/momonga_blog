package handler

import (
	"context"
	"errors"
	"momonga_blog/api"
	"momonga_blog/internal/types"
	"momonga_blog/internal/upload"
)

func (h *Handler) UploadImage(ctx context.Context, req *api.UploadImageReq) (api.UploadImageRes, error) {
	image := req.Image
	// 10MBまで
	if image.Size > 1024*1024*10 {
		return h.NewBadRequest(ctx, "image size is too big", nil), nil
	}
	// 拡張子チェック
	if types.IsAllowedExtension(image.Header.Get("Content-Type")) {
		err := errors.New("image extension is not allowed")
		return h.NewBadRequest(ctx, "image extension is not allowed", err), nil
	}

	// 画像を保存
	useCase := upload.NewImageUseCase()
	_, err := useCase.UploadImage(image.File, image.Name)
	if err != nil {
		return h.NewBadRequest(ctx, "failed to upload image", err), nil
	}

	return &api.NotContent{
		Status: 204,
		Data:   api.NotContentData{},
		Error: api.NotContentError{},
	}, nil
}