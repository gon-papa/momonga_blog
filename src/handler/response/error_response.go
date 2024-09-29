package response

import (
	"momonga_blog/api"
	"momonga_blog/config"
	"momonga_blog/consts"
	"net/http"
)


func ErrorResponse(status int, message string, err error) *api.ErrorResponseStatusCode {
	cnf, configErr := config.GetConfig()
	if configErr != nil {
		return &api.ErrorResponseStatusCode{
			StatusCode: http.StatusInternalServerError,
			Response: api.ErrorResponse{
				Status: http.StatusInternalServerError,
				Data:   api.ErrorResponseData{},
				Error: api.ErrorResponseError{
					Message: api.NewOptString(message),
				},
			},
		}
	}
	if cnf.Env == consts.ProdEnv {
		message = http.StatusText(http.StatusInternalServerError)
	} else {
		if err != nil {
			message = message + " " + err.Error()
		}
	}

	return &api.ErrorResponseStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.ErrorResponse{
			Status: http.StatusInternalServerError,
			Data:   api.ErrorResponseData{},
			Error: api.ErrorResponseError{
				Message: api.NewOptString(message),
			},
		},
	}
}