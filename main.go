package main

import (
	"fmt"

	"github.com/Amane-Fujiwara11/SimpleAPIServer/calc"
	"github.com/Amane-Fujiwara11/SimpleAPIServer/concurrent"
	"github.com/Amane-Fujiwara11/SimpleAPIServer/handler"
	"github.com/Amane-Fujiwara11/SimpleAPIServer/httpclient"
)

func main() {
	port := "8080"

	// まずは急に階乗計算を行う
	number := 10
	factorial, err := calc.Fact(number) // 5の階乗を計算
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%dの階乗は%dです。\n", number, factorial)
	}

	// 並行処理のタスクを実行
	concurrent.RunConcurrentTasks(5) // タスク数を指定

	// HTTPリクエストの取得
	url := "https://jsonplaceholder.typicode.com/posts/1" // テスト用のURL
	httpclient.FetchURL(url)                              // URLを取得

	fmt.Printf("Server is running on http://localhost:%s\n", port)

	handler.GlobalErrorHandler()
}
