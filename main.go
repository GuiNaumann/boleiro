package main

import (
	"boleiro/infrastructure"
	configs "boleiro/settings"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	settings, err := configs.Setup()
	if err != nil {
		log.Println("[main] Error configs.Setup", err)
		return
	}
	serverDomain := settings.GetDomain()

	log.Println(settings)

	router := mux.NewRouter()
	err = infrastructure.Setup(*settings, router)
	if err != nil {
		log.Println("[main] Error infrastructure.Setup", err)
		return
	}

	server := &http.Server{
		Addr:         serverDomain,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      router,
	}

	log.Println("[main] Server is running on", serverDomain)
	log.Fatal(server.ListenAndServe())
}
