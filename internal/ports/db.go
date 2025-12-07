package ports

import "github.com/skyespirates/go-hex/internal/domain"

type DBPort interface {
	Get(id int) (domain.Todo, error)
	GetAll() ([]*domain.Todo, error)
	Save(title string) error
	Update(domain.Todo) error
	Delete(id int) error
}
