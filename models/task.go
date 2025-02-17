package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

const (
	ToDo = iota + 1
	InProgress
	Done
)
