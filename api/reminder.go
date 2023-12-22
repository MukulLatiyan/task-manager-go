package api

import (
	"encoding/json"
	"net/http"

	"github.com/MukulLatiyan/task-manager-go/db"
	"github.com/MukulLatiyan/task-manager-go/models"
	"github.com/go-chi/chi"
)

func CreateNewReminder(w http.ResponseWriter, r *http.Request) {
	var reminder models.Reminder
	err := json.NewDecoder(r.Body).Decode(&reminder)
	if err != nil {
		http.Error(w, "Invalid reminder data", http.StatusBadRequest)
		return
	}

	err = db.CreateReminder(&reminder)
	if err != nil {
		http.Error(w, "Failed to create reminder", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Reminder created successfully"))
}

func GetAllReminders(w http.ResponseWriter, r *http.Request) {
	reminders, err := db.GetAllReminders()
	if err != nil {
		http.Error(w, "Failed to retrieve reminders", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(reminders)
	if err != nil {
		http.Error(w, "Error while Encoding all the reminders", http.StatusInternalServerError)
		return
	}
}

func GetSingleReminderByID(w http.ResponseWriter, r *http.Request) {
	reminderID := chi.URLParam(r, "reminderID")

	reminder, err := db.GetReminderByID(reminderID)
	if err != nil {
		http.Error(w, "Failed to retrieve reminder", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(reminder)
	if err != nil {
		http.Error(w, "Error while Encoding the single reminder", http.StatusInternalServerError)
		return
	}
}

func UpdateReminderByID(w http.ResponseWriter, r *http.Request) {
	reminderID := chi.URLParam(r, "reminderID")

	var updatedReminder models.Reminder
	err := json.NewDecoder(r.Body).Decode(&updatedReminder)
	if err != nil {
		http.Error(w, "Invalid reminder data", http.StatusBadRequest)
		return
	}

	err = db.UpdateReminderByID(reminderID, &updatedReminder)
	if err != nil {
		http.Error(w, "Failed to update reminder", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reminder updated successfully."))
}

func DeleteReminderByID(w http.ResponseWriter, r *http.Request) {
	reminderID := chi.URLParam(r, "reminderID")

	err := db.DeleteReminderByID(reminderID)
	if err != nil {
		http.Error(w, "Failed to delete reminder", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reminder deleted successfully."))
}
