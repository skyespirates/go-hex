package ports

import "github.com/skyespirates/go-hex/internal/domain"

type TodoRepository interface {
	Create(todo *domain.Todo) error
	GetById(id string) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id string) error
	List() ([]*domain.Todo, error)
}
