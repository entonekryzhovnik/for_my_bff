package repository

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) CreateTask() error {
	return nil
}

func (r *TaskRepository) GetTaskByID() (string, error) {
	return "mock_task", nil
}
