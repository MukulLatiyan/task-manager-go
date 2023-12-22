package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"task_manager/api"
	"task_manager/db"
	"task_manager/middleware"
)

func main() {
	db.Connect()

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Group(func(r chi.Router) {
		r.Use(middleware.Authentication)

		r.Post("/tasks", api.CreateNewTask)
		r.Get("/tasks", api.GetAllTheTasks)
		r.Get("/tasks/{taskID}", api.GetIndividualTasks)
		r.Put("/tasks/{taskID}", api.UpdateTaskByTaskID)
		r.Delete("/tasks/{taskID}", api.DeleteTaskByTaskID)
	})

	http.ListenAndServe(":8080", router)
}
