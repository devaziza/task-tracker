package usecase

import (
	"errors"
	"example/Task-Tracker/models"
	"example/Task-Tracker/repository"
	"time"
)

type TaskUsecase struct {
	Repo *repository.TaskRepository
}

func (u *TaskUsecase) AddTask(desc string) (int, error) {
	tasks, _ := u.Repo.Load()
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}
	task := models.Task{
		ID: newID, Description: desc, Status: models.StatusTodo,
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}
	tasks = append(tasks, task)
	return newID, u.Repo.Save(tasks)
}

func (u *TaskUsecase) UpdateTask(id int, desc string) error {
	tasks, _ := u.Repo.Load()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = desc
			tasks[i].UpdatedAt = time.Now()
			return u.Repo.Save(tasks)
		}
	}
	return errors.New("task tapylmady")
}

func (u *TaskUsecase) DeleteTask(id int) error {
	tasks, _ := u.Repo.Load()
	newTasks := []models.Task{}
	found := false
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else { found = true }
	}
	if !found { return errors.New("task tapylmady") }
	return u.Repo.Save(newTasks)
}

func (u *TaskUsecase) MarkStatus(id int, status string) error {
	tasks, _ := u.Repo.Load()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			return u.Repo.Save(tasks)
		}
	}
	return errors.New("task tapylmady")
}

func (u *TaskUsecase) ListTasks(status string) ([]models.Task, error) {
	tasks, err := u.Repo.Load()
	if status == "" { return tasks, err }
	filtered := []models.Task{}
	for _, t := range tasks {
		if t.Status == status { filtered = append(filtered, t) }
	}
	return filtered, nil
}