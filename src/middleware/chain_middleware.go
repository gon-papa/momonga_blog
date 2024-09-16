package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// ミドルウェアを連結して返す。可読性を高めるために使用する。
// 例: ChainMiddleware(h, mw1, mw2, mw3)
func ChainMiddleware(h http.Handler, ms ...Middleware) http.Handler {
	if len(ms) < 1 {
		return h
	}
	wrapped := h
	for i := len(ms) - 1; i >= 0; i-- {
		wrapped = ms[i](wrapped)
	}
	return wrapped
}