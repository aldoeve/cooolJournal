package main

import (
	"cooolJournal/backend/routes"
	"fmt"
	"log"
	"net/http"
)

const (
	portNumberString = ":8080"
)

func main() {
	// Api calls go here.
	patterns := []string{"/api/saveProfile"}
	funcs := []func(http.ResponseWriter, *http.Request){routes.SaveProfile}
	if len(patterns) != len(funcs) {
		log.Panic("API pattern to functions mismatch")
	}
	for index, apiCall := range patterns {
		http.HandleFunc(apiCall, funcs[index])
	}

	// Pointing the files to the complied vue files.
	server := http.FileServer(http.Dir("./../../frontend/dist"))
	http.Handle("/", server)

	// Starting the server.
	fmt.Println("Server listening on", portNumberString)
	// log.Panic allows for defered functions to still run and allow some recovery
	log.Panic(
		http.ListenAndServe(portNumberString, nil),
	)
}
