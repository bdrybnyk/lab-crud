package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockGuitarRepo struct{ memory map[string]Guitar }

func (m *MockGuitarRepo) Create(g Guitar) error                          { m.memory[g.ID] = g; return nil }
func (m *MockGuitarRepo) GetAll() ([]Guitar, error)                      { return nil, nil }
func (m *MockGuitarRepo) GetByID(id string) (Guitar, error)              { return m.memory[id], nil }
func (m *MockGuitarRepo) Update(id string, g Guitar) error               { return nil }
func (m *MockGuitarRepo) PatchElectric(id string, isElectric bool) error { return nil }
func (m *MockGuitarRepo) Delete(id string) error                         { delete(m.memory, id); return nil }

func TestCreateGuitar_Valid(t *testing.T) {
	mockRepo := &MockGuitarRepo{memory: make(map[string]Guitar)}
	handler := NewGuitarHandler(mockRepo)

	body := []byte(`{"brand":"Fender", "model":"Strat", "strings_count":6}`)
	req := httptest.NewRequest("POST", "/guitars", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.CreateGuitar(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Очікуваний статус - 201, отримано - %d", w.Code)
	}
}

func TestCreateGuitar_InvalidValidation(t *testing.T) {
	mockRepo := &MockGuitarRepo{memory: make(map[string]Guitar)}
	handler := NewGuitarHandler(mockRepo)
	body := []byte(`{"brand":"Fender", "model":"Bass", "strings_count":2}`)
	req := httptest.NewRequest("POST", "/guitars", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.CreateGuitar(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Очікуваний статус - 400 через провал валідації, отримано - %d", w.Code)
	}
}
