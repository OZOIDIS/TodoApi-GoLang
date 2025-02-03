package service

//The service layer is responsible for implementing the business logic of the application.
//It is also the layer that interacts with the repository layer to fetch and store data.

import (
	"github.com/ozoidis/todoapi-golang/models"
	"github.com/ozoidis/todoapi-golang/repository"
)

type TaskService interface {
	GetAll() ([]models.Task, error)
	GetByID(id int) (models.Task, error)
}

type taskService struct {
	taskRepository repository.TaskRepository
}

//Don't need this anymore
/* type taskHandler struct {
	taskService TaskService
} */

func NewTaskService(taskRepository repository.TaskRepository) *taskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

func (s *taskService) GetAll() ([]models.Task, error) {
	tasks, error := s.taskRepository.GetAll()

	if error != nil {
		return nil, error
	}

	return tasks, nil
}

func (s *taskService) GetByID(id int) (models.Task, error) {
	return s.taskRepository.GetByID(id)
}

//Don't need this anymore
/* func NewTaskHandler(taskService TaskService) *taskHandler {
	return &taskHandler{
		taskService: taskService,
	}
} */
