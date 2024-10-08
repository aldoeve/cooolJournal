package main

import (
	"fmt"
	"log"
	"net/http"
	"cooolJournal/backend/setup"
)

func main() {
	//Pointing the files to the complied vue files.
	server := http.FileServer(http.Dir("./../../frontend/dist"))
	http.Handle("/", server)

	//Handle Setup
	flags := make(map[string]bool)
	setup.SetupFlags(flags)
	setup.SetupDefault(flags)

	//Starting the server.
	fmt.Println("Server listening on :8080")
	log.Panic(
		http.ListenAndServe(":8080", nil),
	)

	

}
