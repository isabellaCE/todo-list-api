package controllers

import (
	"api/entities"
	"api/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type TaskController struct {
	Repo *repository.TaskRepository
}

func NewTaskController(repo *repository.TaskRepository) *TaskController {
	return &TaskController{Repo: repo}
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entities.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := c.Repo.CreateTask(&task); err != nil {
		fmt.Println("Erro ao criar tarefa:", err)
		http.Error(w, "Erro ao criar tarefa", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
