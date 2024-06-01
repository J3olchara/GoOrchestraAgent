package workers

import (
	"bytes"
	"encoding/json"
	"github.com/J3olchara/GoOrchestraAgnet/app/tasks"
	"log"
	"net/http"
	"os"
	"time"
)

type Worker struct {
	ID       int
	client   *http.Client
	timeout  time.Duration
	FetchUrl string
	out      chan struct{}
}

func NewWorker(id int, timeout time.Duration, FetchUrl string, out chan struct{}) *Worker {
	return &Worker{
		ID:       id,
		client:   &http.Client{},
		timeout:  timeout,
		FetchUrl: FetchUrl,
		out:      out,
	}
}

func (w *Worker) Work() {
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				task := w.GetTask()
				if task == nil {
					continue
				}
				log.Printf("Worker %d got Task %d\n", w.ID, task.ID)
				data, err := json.Marshal(map[string]interface{}{
					"id":     task.ID,
					"result": task.Do(),
				})
				if err != nil {
					log.Println(err)
					return
				}
				req, _ := http.NewRequest(http.MethodPost, w.FetchUrl, bytes.NewReader(data))
				_, err = w.client.Do(req)
				if err != nil {
					log.Println(err)
					return
				}
			case <-w.out:
				return
			}
		}
	}()
}

func (w *Worker) GetTask() *tasks.Task {
	var task map[string]tasks.Task
	resp, err := http.Get(os.Getenv("ORCHESTRA_URL") + os.Getenv("TASK_FETCH_URL"))
	if err != nil {
		log.Println(err)
		return nil
	}
	if resp.StatusCode != 200 {
		return nil
	}
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		log.Println(err)
		return nil
	}
	ret := task["task"]
	return &ret
}
