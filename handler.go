package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type GuitarHandler struct {
	repo     GuitarRepository
	validate *validator.Validate
}

func NewGuitarHandler(r GuitarRepository) *GuitarHandler {
	return &GuitarHandler{repo: r, validate: validator.New()}
}

// CREATE
func (h *GuitarHandler) CreateGuitar(w http.ResponseWriter, r *http.Request) {
	var g Guitar
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, "Кривий JSON", http.StatusBadRequest)
		return
	}
	if g.ID == "" {
		g.ID = uuid.NewString()
	}

	if err := h.validate.Struct(g); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(g); err != nil {
		http.Error(w, "Помилка БД", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(g)
}

// GET ALL
func (h *GuitarHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	guitars, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Помилка БД", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(guitars)
}

// GET BY ID
func (h *GuitarHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	g, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "Не знайдено", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(g)
}

// PUT (Full Update)
func (h *GuitarHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var g Guitar
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, "Помилка JSON", http.StatusBadRequest)
		return
	}
	g.ID = id
	if err := h.validate.Struct(g); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.repo.Update(id, g)
	w.WriteHeader(http.StatusNoContent)
}

// PATCH (Partial Update)
func (h *GuitarHandler) PatchElectric(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var data struct {
		IsElectric bool `json:"is_electric"`
	}
	json.NewDecoder(r.Body).Decode(&data)
	h.repo.PatchElectric(id, data.IsElectric)
	w.WriteHeader(http.StatusNoContent)
}

// DELETE
func (h *GuitarHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	h.repo.Delete(id)
	w.WriteHeader(http.StatusNoContent)
}
