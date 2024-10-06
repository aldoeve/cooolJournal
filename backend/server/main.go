package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	portNumberString = ":8080"
)

func main() {
	//Pointing the files to the complied vue files.
	server := http.FileServer(http.Dir("./../../frontend/dist"))
	http.Handle("/", server)

	//Starting the server.
	fmt.Println("Server listening on", portNumberString)
	//log.Panic allows for defered functions to still run and allow some recovery
	log.Panic(
		http.ListenAndServe(portNumberString, nil),
	)
}
