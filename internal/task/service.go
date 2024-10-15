package task

type Service interface {
	GetAll() []Task
	GetByID(id int) (Task, error)
	Create(task Task) Task
	Update(id int, task Task) (Task, error)
	Delete(id int) error
}

type taskService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &taskService{repo: repo}
}

func (s *taskService) GetAll() []Task {
	return s.repo.GetAll()
}

func (s *taskService) GetByID(id int) (Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) Create(task Task) Task {
	return s.repo.Create(task)
}

func (s *taskService) Update(id int, task Task) (Task, error) {
	return s.repo.Update(id, task)
}

func (s *taskService) Delete(id int) error {
	return s.repo.Delete(id)
}
