package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// home handler function which writes a byte slice containing
// "Hello from tsnippet" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if current request URL path exactly mathces "/".
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from tsnippet"))
}

// snippetView handler function which displays a specific snippet.
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of id parameter from the query string
	// convert it to an integer and check for errors.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

// snippetCreate handler function which creates a new snippet.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Check wether the request method is POST.
	if r.Method != "POST" {
		// If it's not send a 405 status code in response.
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
