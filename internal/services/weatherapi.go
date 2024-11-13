package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperaturaFromWeatherAPI(cidade string) (float64, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	resp, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, url.QueryEscape(cidade)))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	return data.Current.TempC, nil
}
