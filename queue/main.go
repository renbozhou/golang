package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	//ch := make(chan string, 10)
	//wgPd := new(sync.WaitGroup)
	//wgCs := new(sync.WaitGroup)
	//for i := 0; i < 3; i++ {
	//	wgPd.Add(1)
	//	go producer(i, wgPd, ch)
	//}
	//
	//for i := 0; i < 3; i++ {
	//	wgCs.Add(1)
	//	go consumer(i, wgCs, ch)
	//}
	//
	//wgPd.Wait()
	//close(ch)
	//wgCs.Wait()
	//http.put("http://localhost:8500/v1/agent/service/deregister/gateway-192-168-43-22-8082")

	//resp, err := http.Get("http://localhost:8500/v1/health/state/critical")

	req, _ := http.NewRequest("PUT", "http://localhost:8500/v1/agent/service/deregister/tech-192-168-43-22-19021", nil)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req) // http.Get("http://localhost:8500/v1/health/state/critical")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
