package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marcosocram/fullcycle-weather-api/internal/app"
	"github.com/marcosocram/fullcycle-weather-api/internal/services"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// Função de inicialização para o teste
func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/weather", app.GetWeatherHandler).Methods("GET")
	return r
}

func TestGetWeatherHandler_Success(t *testing.T) {
	tempC := 25.0

	serviceViaCEP := &MockViaCEPService{}
	// Mock da função GetCidadeFromCEP
	serviceViaCEP.On("GetCidadeFromCEP", "01001000").Return("São Paulo", nil)
	serviceViaCEP.GetCidadeFromCEP("01001000")
	serviceViaCEP.AssertExpectations(t)

	// Mock da função GetTemperaturaFromWeatherAPI
	serviceWeatherAPI := &MockWeatherAPIService{}
	serviceWeatherAPI.On("GetTemperaturaFromWeatherAPI", "São Paulo").Return(tempC, nil)
	serviceWeatherAPI.GetTemperaturaFromWeatherAPI("São Paulo")
	serviceWeatherAPI.AssertExpectations(t)

	req, _ := http.NewRequest("GET", "/weather?cep=01001000", nil)
	rr := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]float64
	json.Unmarshal(rr.Body.Bytes(), &response)
	//
	assert.InDelta(t, 25.0, response["temp_C"], 10.1)
	assert.InDelta(t, 77.0, response["temp_F"], 20.1)
	assert.InDelta(t, 298.15, response["temp_K"], 30.1)
}

func TestGetWeatherHandler_InvalidZipcode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/weather?cep=1234567", nil) // CEP com menos de 8 dígitos
	rr := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, "invalid zipcode\n", rr.Body.String())
}

func TestGetWeatherHandler_ZipcodeNotFound(t *testing.T) {
	serviceViaCEP := &MockViaCEPService{}
	// Mock da função GetCidadeFromCEP para retornar um erro de CEP não encontrado
	serviceViaCEP.On("GetCidadeFromCEP", "99999999").Return("", services.ErrCEPNotFound)

	req, _ := http.NewRequest("GET", "/weather?cep=99999999", nil) // CEP inexistente
	rr := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "can not find zipcode\n", rr.Body.String())
}
