package repository

import (
	"encoding/json"
	"os"
	"example/Task-Tracker/models"
)

type TaskRepository struct {
	FilePath string
}

func NewTaskRepository(path string) *TaskRepository {
	return &TaskRepository{FilePath: path}
}

func (r *TaskRepository) Load() ([]models.Task, error) {
	file, err := os.ReadFile(r.FilePath)
	if err != nil {
		if os.IsNotExist(err) { return []models.Task{}, nil }
		return nil, err
	}
	var tasks []models.Task
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func (r *TaskRepository) Save(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil { return err }
	return os.WriteFile(r.FilePath, data, 0644)
}