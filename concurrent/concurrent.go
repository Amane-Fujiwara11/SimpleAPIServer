package concurrent

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// タスクの実行をシミュレートする関数
func performTask(id int, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	waitTime := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(waitTime)

	ch <- fmt.Sprintf("Task %d completed after %v", id, waitTime)
}

// 並行タスクを実行する関数
func RunConcurrentTasks(numTasks int) {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	ch := make(chan string, numTasks)

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go performTask(i+1, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// チャネルから結果を受け取る
	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("All tasks completed.")
}
