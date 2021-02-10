package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

// helloHandler responds to requests by saying hi to name
func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	log.Printf("Issued greeting to %s", name)
	fmt.Fprintf(w, "Hello, %s", name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Queries("name", "{name}").Methods("GET")
	// r.HandleFunc("/", indexHandler).Methods("GET")

	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
