package mysql

import (
	"database/sql"
	"errors"

	"github.com/skyespirates/go-hex/internal/domain"
	"github.com/skyespirates/go-hex/internal/usecases"
)

func (a *Adapter) Create(todo *domain.Todo) error {
	query := `INSERT INTO todos (id, title, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	args := []interface{}{todo.Id, todo.Title, todo.Completed, todo.CreatedAt, todo.UpdatedAt}
	_, err := a.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (a *Adapter) GetById(id string) (*domain.Todo, error) {
	var todo domain.Todo
	query := `SELECT id, title, completed, created_at, updated_at FROM todos WHERE id = ?`
	err := a.db.QueryRow(query, id).Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, usecases.ErrNotFound
		default:
			return nil, err
		}
	}
	return &todo, err
}

func (a *Adapter) Update(todo *domain.Todo) error {
	query := `UPDATE todos SET title = ?, completed = ?, updated_at = ? WHERE id = ?`
	args := []interface{}{todo.Title, todo.Completed, todo.UpdatedAt, todo.Id}

	result, err := a.db.Exec(query, args...)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return usecases.ErrNotFound
	}

	return nil
}

func (a *Adapter) Delete(id string) error {
	query := `DELETE FROM todos WHERE id = ?`
	result, err := a.db.Exec(query, id)
	if err != nil {
		return nil
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (a *Adapter) List() ([]*domain.Todo, error) {
	var todos []*domain.Todo

	query := `SELECT id, title, completed, created_at, updated_at FROM todos`

	rows, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var t domain.Todo

		err := rows.Scan(&t.Id, &t.Title, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, &t)
	}

	return todos, nil
}
