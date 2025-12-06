package task2

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全

func mutexCounter() {
	var lock sync.Mutex

	var counter = 0
	for i := 0; i < 10; i++ {
		go func() {
			lock.Lock()
			for j := 0; j < 1000; j++ {
				counter++
			}
			lock.Unlock()
		}()
	}

	time.Sleep(time.Second * 4)
	fmt.Printf("共享计数器counter: %v\n", counter)
}

func atomicCounter() {
	ai := atomic.Int64{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				ai.Add(1)
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("无锁计数器", ai.Load())
}

func Lock() {
	mutexCounter()

	atomicCounter()
}
