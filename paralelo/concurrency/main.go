package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	requestId := make(chan int)
	concurrency := 5
	for i := 1; i <= concurrency; i++ {
		go worker(requestId, i)
	}
	for i := 0; i < 20; i++ {
		requestId <- i
	}
}
func worker(requestId chan int, worker int) {
	for r := range requestId {
		res, err := http.Get("http://localhost:8585/product")
		if err != nil {
			log.Fatal("Error")
		}
		defer res.Body.Close()
		//content, _ := ioutil.ReadAll(res.Body)
		//fmt.Println("Worker %d. RequestId: %d. Content: %s",worker, r, string(content))
		fmt.Println("Worker %d. RequestId: %d.\n", worker, r)
		r := rand.Intn(10)
		time.Sleep(time.Duration(r) * time.Second)
	}
}
