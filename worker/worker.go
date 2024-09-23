package worker

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"luigimorel.com/orchestra/task"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("unimplemented")
}

func (w *Worker) RunTask() {
	fmt.Println("unimplemented")
}

func (w *Worker) StopTask() {
	fmt.Println("unimplemented")
}

func (w *Worker) StartTask() {
	fmt.Println("unimplemented")
}
