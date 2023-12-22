package main

import (
	"github.com/MukulLatiyan/task-manager-go/api"
	"github.com/MukulLatiyan/task-manager-go/db"
	"github.com/MukulLatiyan/task-manager-go/middleware"
	"github.com/go-chi/chi"
	"net/http"
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

		r.Post("/reminders", api.CreateNewReminder)
		r.Get("/reminders", api.GetAllReminders)
		r.Get("/reminders/{reminderID}", api.GetSingleReminderByID)
		r.Put("/reminders/{reminderID}", api.UpdateReminderByID)
		r.Delete("/reminders/{reminderID}", api.DeleteReminderByID)
	})

	http.ListenAndServe(":8080", router)
}
