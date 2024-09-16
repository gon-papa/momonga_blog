package router

import (
	"encoding/json"
	"momonga_blog/middleware"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", middleware.ChainMiddleware(http.HandlerFunc(resource1ListHandler)))

	return mux
}

func resource1ListHandler(w http.ResponseWriter, r *http.Request) {
    // 返すデータを準備
    data := map[string]string{
        "message": "Resource 11",
    }

    // Content-Typeヘッダーを設定
    w.Header().Set("Content-Type", "application/json")

    // データをJSONにエンコードしてレスポンスに書き込む
    if err := json.NewEncoder(w).Encode(data); err != nil {
        // エラーが発生した場合は500エラーを返す
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}