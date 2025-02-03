package tests

import (
	"testing"

	"github.com/ozoidis/todoapi-golang/models"
	"github.com/ozoidis/todoapi-golang/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) GetAll() ([]models.Task, error) {
	args := m.Called()
	return args.Get(0).([]models.Task), args.Error(0)
}

func (m *mockRepo) GetByID(i int) (models.Task, error) {
	args := m.Called()
	return args.Get(0).(models.Task), args.Error(1)
}

func Test_GetAll_Should_Return_Nil_When_Error_Found(t *testing.T) {
	//arrange
	m := new(mockRepo)
	s := service.NewTaskService(m)
	expectedTasks := []models.Task{}

	m.On("GetAll", 1).Return(expectedTasks, nil)

	//act
	ts, er := s.GetAll()

	//assert
	assert.Error(t, er)
	assert.Nil(t, ts)
}
