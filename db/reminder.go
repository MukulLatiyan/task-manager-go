package db

import (
	"github.com/MukulLatiyan/task-manager-go/models"
)

func GetAllReminders() ([]models.Reminder, error) {
	var reminders []models.Reminder
	err := DB.Model(&reminders).Select()
	return reminders, err
}

func GetReminderByID(reminderID string) (models.Reminder, error) {
	var reminder models.Reminder
	err := DB.Model(&reminder).Where("id = ?", reminderID).Select()
	return reminder, err
}

func CreateReminder(reminder *models.Reminder) error {
	_, err := DB.Model(reminder).Insert()
	return err
}

func UpdateReminderByID(reminderID string, updatedReminder *models.Reminder) error {
	_, err := DB.Model(updatedReminder).Where("id = ?", reminderID).Update()
	return err
}

func DeleteReminderByID(reminderID string) error {
	_, err := DB.Model(&models.Reminder{}).Where("id = ?", reminderID).Delete()
	return err
}
