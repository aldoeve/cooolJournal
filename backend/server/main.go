package main

import (
	"cooolJournal/backend/routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	const (
		portNumberString = ":8080"
		vueBianaryPath   = "./../../frontend/dist"
	)

	// Api calls go here.
	patterns := []string{
		"/api/saveProfile", "/api/loginUser",
	}
	funcs := []func(http.ResponseWriter, *http.Request){
		routes.SaveProfile, routes.LoginUser,
	}

	if len(patterns) != len(funcs) {
		log.Panic("API pattern to functions mismatch")
	}
	for index, apiCall := range patterns {
		http.HandleFunc(apiCall, funcs[index])
	}

	// Pointing the files to the complied vue files.
	server := http.FileServer(http.Dir(vueBianaryPath))

	// Serves the index.html when client side reloads page if needed.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := vueBianaryPath + r.URL.Path
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			http.ServeFile(w, r, "./"+vueBianaryPath+"/index.html")
			return
		} else if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		server.ServeHTTP(w, r)
	})

	// Starting the server.
	fmt.Println("Server listening on", portNumberString)
	// log.Panic allows for defered functions to still run and allow some recovery
	log.Panic(
		http.ListenAndServe(portNumberString, nil),
	)
}
