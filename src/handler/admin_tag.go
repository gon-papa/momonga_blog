package handler

import (
	"context"
	"momonga_blog/api"
	"momonga_blog/internal/tag"
	"momonga_blog/internal/types"
	"momonga_blog/repository/model"
	"net/http"
)

func (h *Handler) CreateTag(ctx context.Context, req *api.TagCreateRequest) (api.CreateTagRes, error) {
	useCase := tag.NewTagUseCase()
	err := useCase.CreateTag(req.Name)
	if err != nil {
		return h.NewBadRequest(ctx, "failed to create tag", err), nil
	}

	return &api.NotContent{
		Status: http.StatusOK,
		Data: api.NotContentData{},
		Error: api.NotContentError{},
	}, nil
}


func (h *Handler) GetTagList(ctx context.Context, params api.GetTagListParams) (api.GetTagListRes, error) {
	useCase := tag.NewTagUseCase()
	var tags []*model.Tag
	var err error

	if params.UUID.IsSet() {
		blogUuid := types.NewUuid(params.UUID.Value.String())
		tags, err = useCase.GetTagList(&blogUuid)
		if err != nil {
			return h.NewBadRequest(ctx, "failed to get tag list", err), nil
		}
	} else {
		tags, err = useCase.GetTagList(nil)
		if err != nil {
			return h.NewBadRequest(ctx, "failed to get tag list", err), nil
		}
	}

	resTags := types.Map(tags, func(tag *model.Tag) api.Tag {
		return api.Tag{
			UUID: api.NewOptString(tag.UUID),
			Name: api.NewOptString(tag.Name),
		}
	})


	return &api.TagListResponse{
		Status: http.StatusOK,
		Data: api.TagListResponseData{
			Tags: resTags,
		},
		Error: api.TagListResponseError{},
	}, nil
}