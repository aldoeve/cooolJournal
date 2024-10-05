package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Pointing the files to the complied vue files.
	server := http.FileServer(http.Dir("./../../frontend/dist"))
	http.Handle("/", server)

	//Starting the server.
	fmt.Println("Server listening on :8080")
	log.Panic(
		http.ListenAndServe(":8080", nil),
	)

}
