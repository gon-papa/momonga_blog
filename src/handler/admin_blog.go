package handler

import (
	"context"
	"momonga_blog/api"
	r "momonga_blog/handler/resource"
	"momonga_blog/internal/blog"
	t "momonga_blog/internal/types"
	"net/http"
)

func (h *Handler) GetBlogList(ctx context.Context, params api.GetBlogListParams) (api.GetBlogListRes, error) {
	useCase := blog.NewBlogUseCase()
	var page t.Page = t.NewPage(int(params.Page.Value))
	var limit t.Limit = t.NewLimit(int(params.Limit.Value))
	blogs, err := useCase.GetBlogList(page, limit)

	if err != nil {
		return h.NewBadRequest(ctx, "failed to get blog list", err), nil
	}

	blogList := r.MapBlogsToAPI(blogs.Blogs)

	return &api.BlogResponse{
		Status: http.StatusOK,
		Data: api.BlogResponseData{
			Blogs: blogList,
			Pagenation: r.MapPaginationToAPI(int(blogs.Total), blogs.Page, blogs.Limit),
		},
		Error: api.BlogResponseError{},
	}, nil
}



