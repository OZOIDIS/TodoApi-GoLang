package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/ozoidis/todoapi-golang/handlers"
	"github.com/ozoidis/todoapi-golang/repository"
	"github.com/ozoidis/todoapi-golang/service"
)

func main() {

	//Create a new instance of TaskService
	taskRepository := repository.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)

	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", taskHandler.GetAll).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", taskHandler.GetTaskByID).Methods("GET")

	//Server
	server := http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
	}

	//Get task by ID
	task1, err := taskService.GetByID(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(task1)

	server.ListenAndServe()
}
