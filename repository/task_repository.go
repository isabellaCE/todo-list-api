package repository

import (
	"api/entities"
	"database/sql"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) CreateTask(task *entities.Task) error {
	query := "INSERT INTO public.tasks (title, description, completed) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, task.Title, task.Description, task.Completed)
	return err
}
