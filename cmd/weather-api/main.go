package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/marcosocram/fullcycle-weather-api/internal/app"

	"github.com/gorilla/mux"
)

func main() {
	// Desabilitar a verificação do certificado SSL
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	r := mux.NewRouter()
	r.HandleFunc("/weather", app.GetWeatherHandler).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
