package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestGetZipCode(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/v1/temperature/{zipCode}", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("zipCode", "29902555")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := func(w http.ResponseWriter, r *http.Request) {
	}
	handler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
