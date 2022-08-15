package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 8080

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
	}

	fmt.Fprintf(w, "POST Request successful.")
	skyrimOwner := r.FormValue("skyrimOwner")
	ulfricStatus := r.FormValue("ulfricStatus")

	fmt.Fprintf(w, "Skyrim belongs to the: = %s\n", skyrimOwner)
	fmt.Fprintf(w, "Ulfric Stormcloak is a: = %s\n", ulfricStatus)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 - not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Method rejected.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "You there, you're finally awake.")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("Starting server at: ", PORT)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
