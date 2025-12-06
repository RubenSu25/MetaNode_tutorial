package task2

import (
	"fmt"
	"sync"
	"time"
)

func twoGoroutine() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 1; i < 11; i++ {
			if i%2 == 1 {
				fmt.Printf("奇数: %v\n", i)
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 2; i < 11; i++ {
			if i%2 == 0 {
				fmt.Printf("偶数: %v\n", i)
			}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("完毕")
}

func scheduler(tasks []func()) {
	if len(tasks) == 0 {
		return
	}
	for _, task := range tasks {
		go func() {
			start_time := time.Now()
			task()
			fmt.Printf("方法执行时间: %v\n", time.Since(start_time))
		}()
	}
}

func Goroutine() {

	twoGoroutine()

	tasks := []func(){func() {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("任务1")
	}, func() {
		fmt.Println("任务2")
	}}
	scheduler(tasks)
}
