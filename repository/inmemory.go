package repository

import (
	"fmt"

	"github.com/ozoidis/todoapi-golang/models"
)

type TaskRepository interface {
	GetAll() ([]models.Task, error)
	GetByID(id int) (models.Task, error)
}

type InMemoryTaskRepository struct {
	tasks []models.Task
}

// Creates a new instance of InMemoryTaskRepository with dummy data
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: []models.Task{
			{ID: 1, Title: "Task 1", Description: "Description 1", Status: models.ToDo},
			{ID: 2, Title: "Task 2", Description: "Description 2", Status: models.InProgress},
			{ID: 3, Title: "Task 3", Description: "Description 3", Status: models.Done},
		},
	}
}

func (r *InMemoryTaskRepository) GetAll() ([]models.Task, error) {
	return r.tasks, nil
}

func (r *InMemoryTaskRepository) GetByID(id int) (models.Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return models.Task{}, fmt.Errorf("task with ID %d not found", id)
}
