package usecases

import (
	"errors"
	"time"

	"github.com/skyespirates/go-hex/internal/domain"
	"github.com/skyespirates/go-hex/internal/ports"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("todo not found")

type TodoService struct {
	repo ports.TodoRepository
}

func NewTodoService(r ports.TodoRepository) *TodoService {
	return &TodoService{repo: r}
}

func (s *TodoService) Create(title string) (*domain.Todo, error) {
	now := time.Now().UTC()

	t := &domain.Todo{
		Id:        uuid.NewString(),
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(t); err != nil {
		return nil, err
	}

	return t, nil
}

func (s *TodoService) GetById(id string) (*domain.Todo, error) {
	todo, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	if todo == nil {
		return nil, ErrNotFound
	}

	return todo, nil
}

func (s *TodoService) Update(todo *domain.Todo) error {
	todo.UpdatedAt = time.Now()
	err := s.repo.Update(todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoService) Delete(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) List() ([]*domain.Todo, error) {
	todos, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
