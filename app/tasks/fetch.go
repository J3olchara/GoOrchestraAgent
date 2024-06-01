package tasks

import (
	"time"
)

type Task struct {
	ID            int     `json:"id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime int     `json:"operation_time"`
}

func (t *Task) Do() float64 {
	time.Sleep(time.Duration(t.OperationTime) * time.Millisecond)
	if t.Operation == "*" {
		return t.Arg1 * t.Arg2
	}
	if t.Operation == "/" {
		return t.Arg1 / t.Arg2
	}
	if t.Operation == "+" {
		return t.Arg1 + t.Arg2
	}
	if t.Operation == "-" {
		return t.Arg1 - t.Arg2
	}
	return 0
}
