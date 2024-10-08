package middleware

import (
	"encoding/json"
	"momonga_blog/handler/response"
	"momonga_blog/internal/logging"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                // エラーログを記録
                logging.ErrorLogger.Error("Panic occurred", "error", err)
                logging.LogWithStackTrace()

                // エラーレスポンスを構築
                errorResponse := response.ErrorResponse(
                    http.StatusInternalServerError,
                    http.StatusText(http.StatusInternalServerError),
                    nil,
                )

                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(errorResponse.StatusCode)

				// エラーレスポンスをJSONに変換
                responseBody, marshalErr := json.Marshal(errorResponse.Response)
                if marshalErr != nil {
                    logging.ErrorLogger.Error("Failed to marshal error response", "error", marshalErr)
                    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
                    return
                }

                // レスポンスを書き込む
                if _, writeErr := w.Write(responseBody); writeErr != nil {
                    logging.ErrorLogger.Error("Failed to write response", "error", writeErr)
                }
            }
        }()
        next.ServeHTTP(w, r)
    })
}