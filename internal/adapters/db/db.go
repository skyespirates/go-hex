package db

import (
	"database/sql"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(dsn string) (*Adapter, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("invalid dsn, error: %v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &Adapter{db}, err
}

func (a *Adapter) Save(title string) error {
	query := `INSERT INTO todos (title) VALUES (?)`

	_, err := a.db.Exec(query, title)
	return err
}
