package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jeanSagaz/temperature-system-by-cep-gcp/internal/application/services"
)

func HandleRequests() {
	fmt.Println("Rest API v2.0 - chi Routers")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/v1/temperature", Router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Router(r chi.Router) {
	r.Get("/{zipCode}", getZipCode)
}

func getZipCode(w http.ResponseWriter, r *http.Request) {
	zipCode := chi.URLParam(r, "zipCode")

	w.Header().Add("Content-Type", "application/json")
	if len(zipCode) != 8 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	responseViaCep, err := services.GetViaCepApiService(zipCode)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	if len(responseViaCep.Localidade) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	responseWeather, err := services.GetWeatherApiService(responseViaCep.Localidade)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := services.FormatTemperatureService(responseViaCep.Localidade, responseWeather.Current.TempC)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
