package handler

import (
	"context"
	"momonga_blog/api"
	r "momonga_blog/handler/resource"
	"momonga_blog/internal/blog"
	"momonga_blog/internal/logging"
	t "momonga_blog/internal/types"
	"net/http"
)

func (h *Handler) GetBlogList(ctx context.Context, params api.GetBlogListParams) (api.GetBlogListRes, error) {
	useCase := blog.NewBlogUseCase()
	var page t.Page = t.NewPage(params.Page.Value)
	var limit t.Limit = t.NewLimit(params.Limit.Value)
	blogs, err := useCase.GetBlogList(page, limit)

	if err != nil {
		logging.ErrorLogger.Error("Failed to logout", "error", err)
		return h.NewBadRequest(ctx, "failed to get blog list", err), nil
	}

	blogList := r.MapBlogsToAPI(blogs.Blogs)

	return &api.BlogListResponse{
		Status: http.StatusOK,
		Data: api.BlogListResponseData{
			Blogs: blogList,
			Pagenation: r.MapPaginationToAPI(blogs.Total, blogs.Page, blogs.Limit),
		},
		Error: api.BlogListResponseError{},
	}, nil
}

func (h *Handler) GetBlog(ctx context.Context, params api.GetBlogParams) (api.GetBlogRes, error) {
	useCase := blog.NewBlogUseCase()
	uuid := t.NewUuid(params.UUID.String())
	blog, err := useCase.GetBlog(uuid)
	if err != nil {
		logging.ErrorLogger.Error("Failed to logout", "error", err)
		return h.NewBadRequest(ctx, "failed to get blog", err), nil
	}

	rBlog := api.Blog{
		UUID: api.NewOptString(blog.UUID),
		Year: api.NewOptInt(blog.Year),
		Month: api.NewOptInt(blog.Month),
		Day: api.NewOptInt(blog.Day),
		Title: api.NewOptString(blog.Title),
		Body: api.NewOptString(blog.Body),
		IsShow: api.NewOptBool(blog.IsShow),
		CreatedAt: api.NewOptString(blog.CreatedAt.String()),
		UpdatedAt: api.NewOptString(blog.UpdatedAt.String()),
		DeletedAt: api.NewOptString(blog.DeletedAt.String()),
	}

	return &api.BlogResponse{
		Status: http.StatusOK,
		Data: api.BlogResponseData{
			Blog: api.NewOptBlog(rBlog),
		},
		Error: api.BlogResponseError{},
	}, nil
}

func (h *Handler) CreateBlogPost(ctx context.Context, params *api.BlogPostRequest) (api.CreateBlogPostRes, error) {
	blogData := t.NewCreateBlogData(
		nil,
		nil,
		nil,
		params.Title,
		params.Body,
		params.IsShow,
	)

	var tags []t.CreateTagData
	if len(params.Tags) != 0 {
		for _, tag := range params.Tags {
			tagData := t.NewCreateTagData(tag)
			tags = append(tags, tagData)
		}
	} else {
		tags = nil
	}

	useCase := blog.NewBlogUseCase()
	_, err := useCase.CreateBlog(blogData, tags)
	if err != nil {
		logging.ErrorLogger.Error("Failed to logout", "error", err)
		return h.NewBadRequest(ctx, "failed to create blog", err), nil
	}


	return &api.NotContent{
		Status: http.StatusOK,
		Data: api.NotContentData{},
		Error: api.NotContentError{},
	}, nil
}