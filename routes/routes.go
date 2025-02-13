package routes

import (
	"api/controllers"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, taskController *controllers.TaskController) {
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			taskController.CreateTask(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
