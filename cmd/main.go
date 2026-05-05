package main

import (
	"fmt"
	"net/http"
	"example/Task-Tracker/handler"
	"example/Task-Tracker/repository"
	"example/Task-Tracker/usecase"
)

func main() {
	repo := repository.NewTaskRepository("tasks.json")
	uc := &usecase.TaskUsecase{Repo: repo}
	h := &handler.TaskHandler{Usecase: uc}

	// CLI komandalarynyň WEB-däki garşylygy:
	http.HandleFunc("/list", h.ListTasks)    // task-tracker list
	http.HandleFunc("/add", h.AddTask)      // task-tracker add
	http.HandleFunc("/delete", h.DeleteTask) // task-tracker delete

	fmt.Println("Web Serwer başlady: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}