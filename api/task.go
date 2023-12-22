package api

import (
	"encoding/json"
	"github.com/MukulLatiyan/task-manager-go/db"
	"github.com/MukulLatiyan/task-manager-go/models"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = db.CreateTask(&newTask)
	if err != nil {
		http.Error(w, "Error inserting new task into the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task created successfully"))
}

func GetAllTheTasks(w http.ResponseWriter, r *http.Request) {
	allTasks, err := db.GetAllTasks()
	if err != nil {
		http.Error(w, "Error while retrieving all tasks from the database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allTasks)
	if err != nil {
		http.Error(w, "Error while Encoding the tasks", http.StatusInternalServerError)
		return
	}
}

func GetIndividualTasks(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
	if err != nil {
		http.Error(w, "Task ID isn't valid", http.StatusBadRequest)
		return
	}

	singleTask, err := db.GetTaskByID(taskID)
	if err != nil {
		http.Error(w, "Error retrieving task from the database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(singleTask)
	if err != nil {
		http.Error(w, "Error while Encoding the single task", http.StatusInternalServerError)
		return
	}
}

func UpdateTaskByTaskID(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
	if err != nil {
		http.Error(w, "Task ID isn't valid", http.StatusBadRequest)
		return
	}

	var singleTask models.Task
	err = json.NewDecoder(r.Body).Decode(&singleTask)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = db.UpdateTaskByID(taskID, &singleTask)
	if err != nil {
		http.Error(w, "Error updating task in the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task updated successfully."))
}

func DeleteTaskByTaskID(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
	if err != nil {
		http.Error(w, "Task ID isn't valid", http.StatusBadRequest)
		return
	}

	err = db.DeleteTaskByID(taskID)
	if err != nil {
		http.Error(w, "Error while deleting a task from the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted successfully."))
}
