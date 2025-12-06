package task2

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 1; i < 11; i++ {
		ch <- i
	}
}

func receiver(ch chan int) {
	for c := range ch {
		fmt.Printf("通道1: %v\n", c)
	}
}

func Channel() {
	var ch = make(chan int)
	go sender(ch)
	go receiver(ch)

	ch1 := make(chan int, 10)
	go bufferSender(ch1)
	go bufferReceiver(ch1)
	time.Sleep(time.Second)

}

func bufferSender(ch chan<- int) {
	for i := 1; i < 101; i++ {
		ch <- i
	}
	close(ch)
}

func bufferReceiver(ch <-chan int) {
	for c := range ch {
		fmt.Printf("通道2: %v\n", c)
	}
}
