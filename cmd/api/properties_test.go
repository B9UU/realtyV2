package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"realtyV2/internal/data"
	"realtyV2/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
}

func (m *MockStore) GetAll() ([]models.Property, error) {
	return []models.Property{}, nil
}
func (m *MockStore) Search(ctx context.Context, b []string) ([]models.Property, error) {
	return []models.Property{}, nil
}

func (m *MockStore) GetById(id int) (models.Property, error) {
	return models.Property{}, nil
}

func (m *MockStore) AddOne(listing models.Property) error {
	return nil

}
func TestGetProperties(t *testing.T) {
	e := newServer()

	// Mock the application and store
	mockStore := new(MockStore)
	app := &Application{
		store: &data.Store{
			Property: mockStore,
		},
		cache: make(map[string]BoundBox),
	}
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
	}{
		{
			name:     "valid queries",
			urlPath:  "/?q=amesterdam&page=1",
			wantCode: http.StatusOK,
		},
		{
			name:     "valid queries",
			urlPath:  "/?q=amesterdam",
			wantCode: http.StatusOK,
		},
		{
			name:     "invalid queries",
			urlPath:  "/?q=am&page-1",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.urlPath, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if assert.NoError(t, app.GetProperties(c)) {
				assert.Equal(t, tt.wantCode, rec.Result().StatusCode)
				assert.Equal(t, rec.Header().Get("content-Type"), "application/json")
			}
		})
	}
}
