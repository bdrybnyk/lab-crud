package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := InitDB(cfg)
	repo := &PostgresRepo{db: db}
	handler := NewGuitarHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /guitars", handler.CreateGuitar)
	mux.HandleFunc("GET /guitars", handler.GetAll)
	mux.HandleFunc("GET /guitars/{id}", handler.GetByID)
	mux.HandleFunc("PUT /guitars/{id}", handler.Update)
	mux.HandleFunc("PATCH /guitars/{id}", handler.PatchElectric)
	mux.HandleFunc("DELETE /guitars/{id}", handler.Delete)

	fmt.Printf("Сервер працює, порт: %s\n", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(cfg.ServerPort, mux))
}
