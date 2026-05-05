package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"example/Task-Tracker/usecase"
)

type TaskHandler struct {
	Usecase *usecase.TaskUsecase
}

// GET /tasks - Hemme tasklary görmek
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status") // ?status=done görnüşinde süzüp bolýar
	tasks, _ := h.Usecase.ListTasks(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// POST /add?desc=... - Täze task goşmak
func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	desc := r.URL.Query().Get("desc")
	id, err := h.Usecase.AddTask(desc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Task added successfully. ID: " + strconv.Itoa(id)))
}

// DELETE /delete?id=... - Task öçürmek
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	h.Usecase.DeleteTask(id)
	w.Write([]byte("Task deleted successfully"))
}