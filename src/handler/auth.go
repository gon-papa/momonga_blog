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
	token, err := useCase.Login(ctx, req.UserID.Value, req.Password.Value)
	if err != nil {
		logging.ErrorLogger.Error("Failed to login", "error", err)
		return &api.BadRequest{
			Status: api.NewOptInt(http.StatusBadRequest),
			Data: nil,
			Error: api.NewOptBadRequestError(
				api.BadRequestError{
					Message: api.NewOptString(err.Error()),
				},
			),
		}, nil
	}

	data := api.LoginResponseData{
		Token: token.Token,
		RefreshToken: token.RefreshToken,
	}

	return &api.LoginResponse{
		Status: api.NewOptInt(http.StatusOK),
		Data: api.NewOptLoginResponseData(data),
		Error: nil,
	}, nil
}


func (h *Handler) Logout(ctx context.Context) (api.LogoutRes, error) {
	useCase := auth.NewLoginUseCase()
	err := useCase.Logout(ctx)

	if err != nil {
		logging.ErrorLogger.Error("Failed to logout", "error", err)
		return &api.BadRequest{
			Status: api.NewOptInt(http.StatusBadRequest),
			Data: nil,
			Error: api.NewOptBadRequestError(
				api.BadRequestError{
					Message: api.NewOptString(err.Error()),
				},
			),
		}, nil
	}

	return &api.NotContent{
		Status: api.NewOptInt(http.StatusOK),
		Data: nil,
		Error: nil,
	}, nil
}
