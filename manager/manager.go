package manager

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"luigimorel.com/orchestra/task"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkTaskMap   map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("unimplemented")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("unimplemented")
}

func (m *Manager) SendWork() {
	fmt.Println("unimplemented")
}
