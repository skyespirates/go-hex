package inmemory

import (
	"sync"

	"github.com/skyespirates/go-hex/internal/domain"
	"github.com/skyespirates/go-hex/internal/usecases"
)

type InMemoryTodoRepo struct {
	mu    sync.RWMutex
	todos map[string]*domain.Todo
}

func NewTodoRepo() *InMemoryTodoRepo {
	return &InMemoryTodoRepo{
		todos: make(map[string]*domain.Todo),
	}
}

func (r *InMemoryTodoRepo) Create(todo *domain.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.todos[todo.Id] = todo
	return nil
}

func (r *InMemoryTodoRepo) GetById(id string) (*domain.Todo, error) {
	todo, ok := r.todos[id]
	if !ok {
		return nil, usecases.ErrNotFound
	}

	return todo, nil
}

func (r *InMemoryTodoRepo) Update(todo *domain.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.todos[todo.Id]
	if !ok {
		return usecases.ErrNotFound
	}

	r.todos[todo.Id] = todo

	return nil
}

func (r *InMemoryTodoRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.todos[id]
	if !ok {
		return usecases.ErrNotFound
	}

	delete(r.todos, id)
	return nil
}

func (r *InMemoryTodoRepo) List() ([]*domain.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	list := make([]*domain.Todo, 0, len(r.todos))

	for _, todo := range r.todos {
		list = append(list, todo)
	}

	return list, nil
}
