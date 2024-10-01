package handler

import (
	"context"
	"momonga_blog/api"
	"momonga_blog/internal/auth"
	"momonga_blog/internal/logging"
	"net/http"
)



func (h *Handler) Login(ctx context.Context, req *api.LoginRequest) (api.LoginRes, error) {
	useCase := auth.NewLoginUseCase()
	token, err := useCase.Login(ctx, req.UserID, req.Password)
	if err != nil {
		logging.ErrorLogger.Error("Failed to login", "error", err)
		return h.NewBadRequest(ctx, "failed to login", err), nil
	}

	data := api.LoginResponseData{
		Token: token.Token,
		RefreshToken: token.RefreshToken,
	}

	return &api.LoginResponse{
		Status: http.StatusOK,
		Data: api.LoginResponseData{
			Token: data.Token,
			RefreshToken: data.RefreshToken,
		},
		Error: api.LoginResponseError{},
	}, nil
}


func (h *Handler) Logout(ctx context.Context) (api.LogoutRes, error) {
	useCase := auth.NewLoginUseCase()
	err := useCase.Logout(ctx)

	if err != nil {
		logging.ErrorLogger.Error("Failed to logout", "error", err)
		return h.NewBadRequest(ctx, "failed to logout", err), nil
	}

	return &api.NotContent{
		Status: http.StatusOK,
		Data: api.NotContentData{},
		Error: api.NotContentError{},
	}, nil
}

func (h *Handler) RefreshToken(ctx context.Context, req *api.RefreshRequest) (api.RefreshTokenRes, error) {
	useCase := auth.NewLoginUseCase()
	refreshToken := req.RefreshToken
	token, err := useCase.RefreshToken(ctx, refreshToken)
	if err != nil {
		logging.ErrorLogger.Error("Failed to refresh token", "error", err)
		return h.NewBadRequest(ctx, "failed to refresh token", err), nil
	}

	return &api.RefreshResponse{
		Status: http.StatusOK,
		Data: api.RefreshResponseData{
			Token: token.Token,
			RefreshToken: token.RefreshToken,
		},
		Error: api.RefreshResponseError{},
	}, nil
}