package handler

import (
	"context"
	"momonga_blog/api"
	"net/http"
)



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
