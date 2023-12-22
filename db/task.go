package db

import (
	"github.com/MukulLatiyan/task-manager-go/models"
)

func CreateTask(task *models.Task) error {
	_, err := DB.Model(task).Insert()
	return err
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := DB.Model(&tasks).Select()
	return tasks, err
}

func GetTaskByID(taskID int) (models.Task, error) {
	var task models.Task
	err := DB.Model(&task).Where("id = ?", taskID).Select()
	return task, err
}

func UpdateTaskByID(taskID int, updatedTask *models.Task) error {
	_, err := DB.Model(updatedTask).Where("id = ?", taskID).Update()
	return err
}

func DeleteTaskByID(taskID int) error {
	_, err := DB.Model(&models.Task{}).Where("id = ?", taskID).Delete()
	return err
}
