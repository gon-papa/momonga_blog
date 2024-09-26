package handler

import (
	"context"
	"momonga_blog/api"
	"momonga_blog/handler/response"
	"net/http"
)

type Handler struct {}
var _ api.Handler = &Handler{}

func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorResponseStatusCode {
	return response.ErrorResponse(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		err,
	)
}
