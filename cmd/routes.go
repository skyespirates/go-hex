package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	router.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("todo list"))
	})

	router.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("todo list"))
	}).Methods(http.MethodPost)

	return router
}
