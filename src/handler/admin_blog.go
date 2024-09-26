package handler

import (
	"context"
	"momonga_blog/api"
	"net/http"
)

func (h *Handler) GetBlogList(ctx context.Context, params api.GetBlogListParams) (api.GetBlogListRes, error) {
	data := api.GetBlogListOKData{
		Blogs: api.GetBlogListOKDataBlogs{},
		Total: 10,
		Page:  1,
		Limit: 10,
	  }

	return &api.GetBlogListOK{
		Status: api.NewOptInt(http.StatusOK),
		Data: api.NewOptGetBlogListOKData(data),
		Error: nil,
	}, nil
}