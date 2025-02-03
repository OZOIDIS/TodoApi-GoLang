package handlers

//The handler package is responsible for handling the incoming HTTP requests and returning the response to the client.
//It provides the implementation of the HTTP handlers for the API endpoints.

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ozoidis/todoapi-golang/models"
	"github.com/ozoidis/todoapi-golang/service"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]

	idNum, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatalf("Error:", err)
	}

	task, err := h.taskService.GetByID(idNum)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	//Set the Access-Control-Allow-Origin header to allow all origins
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	responseDto := mapModelToDto(task)

	err = json.NewEncoder(w).Encode(responseDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks, error := h.taskService.GetAll()

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	//Set the Access-Control-Allow-Origin header to allow all origins
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	responseDto := make([]models.TaskDto, 0, len(tasks))

	for i := range len(tasks) {
		responseDto = append(responseDto, *mapModelToDto(tasks[i]))
	}

	err := json.NewEncoder(w).Encode(responseDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

}

func mapModelToDto(task models.Task) *models.TaskDto {
	var statusString string = ""

	switch task.Status {
	case 1:
		statusString = "To Do"
	case 2:
		statusString = "In Progress"
	case 3:
		statusString = "Done"
	}

	return &models.TaskDto{
		Title:       task.Title,
		Description: task.Description,
		Status:      statusString,
	}

}
