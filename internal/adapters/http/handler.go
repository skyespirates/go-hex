package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/skyespirates/go-hex/internal/usecases"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *usecases.TodoService
}

func NewHandler(s *usecases.TodoService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", h.homepage).Methods("GET")
	r.HandleFunc("/todos", h.createTodo).Methods("POST")
	r.HandleFunc("/todos", h.getAllTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", h.getTodoById).Methods("GET")
	r.HandleFunc("/todos/{id}", h.updateTodos).Methods("PUT")
	r.HandleFunc("/todos/{id}", h.deleteTodo).Methods("DELETE")
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handler) homepage(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Message string `json:"message"`
	}{
		Message: "Hello, World!",
	}

	writeJSON(w, http.StatusOK, res)
}

type createReq struct {
	Title string `json:"title"`
}

func (h *Handler) createTodo(w http.ResponseWriter, r *http.Request) {
	var input createReq
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	todo, err := h.service.Create(input.Title)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, todo)
}

func (h *Handler) getAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.List()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "failed", 501)
		return
	}

	writeJSON(w, http.StatusOK, todos)
}

func (h *Handler) getTodoById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	todo, err := h.service.GetById(id)
	if err != nil {
		if err == usecases.ErrNotFound {
			http.Error(w, "todo not found", http.StatusNotFound)
			return
		}
		log.Fatal(err)
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

func (h *Handler) updateTodos(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var payload struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	todo, err := h.service.GetById(id)
	if err != nil {
		log.Printf("ada error nich, error %v", err)
		if errors.Is(err, usecases.ErrNotFound) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if payload.Title != "" {
		todo.Title = payload.Title
	}

	todo.Completed = payload.Completed
	todo.UpdatedAt = time.Now().UTC()

	err = h.service.Update(todo)
	if err != nil {
		log.Printf("el bantay, error %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

func (h *Handler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.service.Delete(id)
	if err != nil {
		if err == usecases.ErrNotFound {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
