package main

import (
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the root of this API!")
	fmt.Println("Endpoint Hit: root")
}

func contacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the list of Contacts!")
	fmt.Println("Endpoint Hit: contacts")
}

func handleRequests() {
	http.HandleFunc("/", root)
	http.HandleFunc("/contacts", contacts)
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func main() {
	handleRequests()
}
