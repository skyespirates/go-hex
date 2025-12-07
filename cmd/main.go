package main

import (
	"log"
	"net/http"

	persistance "github.com/skyespirates/go-hex/internal/adapters/db/inmemory"
	handler "github.com/skyespirates/go-hex/internal/adapters/http"
	"github.com/skyespirates/go-hex/internal/usecases"

	"github.com/gorilla/mux"
)

func main() {
	repo := persistance.NewInMemoryRepo()
	svc := usecases.NewTodoService(repo)
	h := handler.NewHandler(svc)

	r := mux.NewRouter()
	h.RegisterRoutes(r)

	addr := ":8080"
	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
