package main

import (
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
	}
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name: ,
		}
	}

	t.Run("successful response", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the handler
		err := app.GetProperties(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, rec.Header().Get("content-Type"), "application/json")
		}

	})
}
