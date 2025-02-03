package models

type TaskDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
