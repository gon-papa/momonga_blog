package handler

import (
	"context"
	"momonga_blog/api"
	"net/http"
)

type Handler struct {}
var _ api.Handler = &Handler{}

func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorResponseStatusCode {
	return &api.ErrorResponseStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.ErrorResponse{
			Status: api.NewOptInt(http.StatusInternalServerError),
			Data:   nil,
			Error:  api.NewOptErrorResponseError(api.ErrorResponseError{
				Message: api.NewOptString("Internal Server Error"),
			}),
		},
	}
}

func (h *Handler) Login(ctx context.Context, req *api.LoginRequest) (api.LoginRes, error) {
	data := api.LoginResponseData{
		Token: "xxxx",
		RefreshToken: "yyyy",
	}

	return &api.LoginResponse{
		Status: api.NewOptInt(http.StatusOK),
		Data: api.NewOptLoginResponseData(data),
		Error: nil,
	}, nil
}


func (h *Handler) Logout(ctx context.Context) (api.LogoutRes, error) {
	return &api.NotContent{
		Status: api.NewOptInt(http.StatusOK),
		Data: nil,
		Error: nil,
	}, nil
}

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
