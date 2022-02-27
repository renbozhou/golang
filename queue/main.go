package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func producer(threadId int, wg *sync.WaitGroup, ch chan<- string) {
	count := 0
	for {
		data := strconv.Itoa(threadId)
		ch <- data
		fmt.Printf("producer: %s\n", data)
		time.Sleep(time.Second * 1)
		count++
		if count > 5 {
			wg.Done()
		}
	}

}
func consumer(threadID int, wg *sync.WaitGroup, ch <-chan string) {

	for v := range ch {
		fmt.Printf("consumer, %s\n", v)
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}
func main() {
	ch := make(chan string, 10)
	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wgPd.Add(1)
		go producer(i, wgPd, ch)
	}

	for i := 0; i < 3; i++ {
		wgCs.Add(1)
		go consumer(i, wgCs, ch)
	}

	wgPd.Wait()
	close(ch)
	wgCs.Wait()

}
