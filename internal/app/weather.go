package app

import (
	"encoding/json"
	"github.com/marcosocram/fullcycle-weather-api/internal/services"
	"github.com/marcosocram/fullcycle-weather-api/pkg"
	"net/http"

	"github.com/pkg/errors"
)

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	if len(cep) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	cidade, err := services.GetCidadeFromCEP(cep)
	if err != nil {
		if errors.Is(err, services.ErrCEPNotFound) {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
		} else {
			http.Error(w, "error retrieving data", http.StatusInternalServerError)
		}
		return
	}

	tempC, err := services.GetTemperaturaFromWeatherAPI(cidade)
	if err != nil {
		http.Error(w, "error retrieving weather data", http.StatusInternalServerError)
		return
	}

	resp := map[string]float64{
		"temp_C": tempC,
		"temp_F": pkg.CelsiusToFahrenheit(tempC),
		"temp_K": pkg.CelsiusToKelvin(tempC),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
