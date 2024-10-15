package task

import (
	"errors"
	"sync"
)

type Repository interface {
	GetAll() []Task
	GetByID(id int) (Task, error)
	Create(task Task) Task
	Update(id int, task Task) (Task, error)
	Delete(id int) error
}

type InMemoryRepository struct {
	sync.Mutex
	tasks  map[int]Task
	nextID int
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

func (r *InMemoryRepository) GetAll() []Task {
	r.Lock()
	defer r.Unlock()
	var taskList []Task
	for _, task := range r.tasks {
		taskList = append(taskList, task)
	}
	return taskList
}

func (r *InMemoryRepository) GetByID(id int) (Task, error) {
	r.Lock()
	defer r.Unlock()
	task, exists := r.tasks[id]
	if !exists {
		return Task{}, errors.New("task not found")
	}
	return task, nil
}

func (r *InMemoryRepository) Create(task Task) Task {
	r.Lock()
	defer r.Unlock()
	task.ID = r.nextID
	r.tasks[task.ID] = task
	r.nextID++
	return task
}

func (r *InMemoryRepository) Update(id int, updatedTask Task) (Task, error) {
	r.Lock()
	defer r.Unlock()
	_, exists := r.tasks[id]
	if !exists {
		return Task{}, errors.New("task not found")
	}
	updatedTask.ID = id
	r.tasks[id] = updatedTask
	return updatedTask, nil
}

func (r *InMemoryRepository) Delete(id int) error {
	r.Lock()
	defer r.Unlock()
	if _, exists := r.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}
