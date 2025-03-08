package service

import "api-gateway/internal/repository"

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask() error {
	// TODO: Вызывать CreateTask у репозитория
	return s.taskRepo.CreateTask()
}

func (s *TaskService) GetTaskByID() (string, error) {
	// TODO: Вызывать GetTaskByID у репозитория
	return s.taskRepo.GetTaskByID()
}
