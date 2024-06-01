package main

import (
	"github.com/J3olchara/GoOrchestraAgnet/app/workers"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Ltime | log.Ldate)
	log.Println("on")
	goroutines, err := strconv.Atoi(os.Getenv("COMPUTING_POWER"))
	if err != nil {
		log.Fatal("environ variable COMPUTING_POWER didn't set")
	}
	out := make(chan struct{})
	timeout, err := strconv.Atoi(os.Getenv("REQUESTS_TIMEOUT_MS"))
	log.Println("started creating workers")
	for i := 1; i <= goroutines; i++ {
		worker := workers.NewWorker(
			i,
			time.Duration(timeout)*time.Millisecond,
			os.Getenv("ORCHESTRA_URL")+os.Getenv("TASK_FETCH_URL"),
			out,
		)
		worker.Work()
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
