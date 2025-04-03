package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// タスクの実行をシミュレートする関数
func performTask(id int, wg *sync.WaitGroup, ch chan<- string) {
	// 関数を実行した後にタスクが完了したことをWaitGroupに通知する定義
	// 関数の最後に実行される
	defer wg.Done()

	// シミュレーションのためにランダムな待機時間
	waitTime := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(waitTime)

	// 完了したタスクの結果をチャネルに送信
	ch <- fmt.Sprintf("Task %d completed after %v", id, waitTime)
}

// 並行タスクを実行する関数
func runConcurrentTasks() {
	rand.Seed(time.Now().UnixNano())

	const numTasks = 5
	// goルーチンを定義
	var wg sync.WaitGroup
	// バッファ付きチャネル
	ch := make(chan string, numTasks)

	// 複数のgoルーチンを起動して並行処理を定義
	for i := 0; i < numTasks; i++ {
		// WaitGroupにタスクを追加；タスク開始
		wg.Add(1)
		go performTask(i+1, &wg, ch)
	}

	go func() {
		wg.Wait() // 全タスクの完了を待つ
		close(ch) // チャネルを閉じる
	}()

	// チャネルから結果を受け取る
	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("All tasks completed.")
}
