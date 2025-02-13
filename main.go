package main

import (
	"api/config"
	"api/controllers"
	"api/repository"
	"api/routes"
	"fmt"
	"net/http"
)

func main() {
	// Conectar ao banco
	db, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Criar repositório e controlador
	taskRepo := repository.NewTaskRepository(db)
	taskController := controllers.NewTaskController(taskRepo)

	// Configurar rotas
	mux := http.NewServeMux()
	routes.SetupRoutes(mux, taskController)

	// Iniciar servidor
	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", mux)
}
