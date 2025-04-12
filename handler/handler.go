package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func errorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// エラーが発生した場合のレスポンスを作成する
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("%v", err)})
			}
		}()
		next.ServeHTTP(w, r) // 次のハンドラーを呼び出す
	})
}

func GlobalErrorHandler() {
	http.Handle("/hello", errorHandler(http.HandlerFunc(helloHandler)))
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 意図的なパニックを発生させる
	panic("something went wrong!!!")
}
