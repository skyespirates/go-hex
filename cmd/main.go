package main

import (
	"log"
	"net/http"

	handler "github.com/skyespirates/go-hex/internal/adapters/http"
	// "github.com/skyespirates/go-hex/internal/adapters/persistances/inmemory"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/skyespirates/go-hex/internal/adapters/persistances/mysql"
	"github.com/skyespirates/go-hex/internal/usecases"
)

func main() {
	// repo := inmemory.NewTodoRepo()

	repo, err := mysql.NewAdapter("root:secret@/mydb?parseTime=true")
	if err != nil {
		return
	}
	svc := usecases.NewTodoService(repo)
	h := handler.NewHandler(svc)

	r := mux.NewRouter()
	h.RegisterRoutes(r)

	addr := ":8080"
	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
