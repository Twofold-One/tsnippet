package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// home handler function which writes a byte slice containing
// "Hello from tsnippet" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if current request URL path exactly mathces "/".
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	// files is a slice containing the pathss to the templates.
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// snippetView handler function which displays a specific snippet.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of id parameter from the query string
	// convert it to an integer and check for errors.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

// snippetCreate handler function which creates a new snippet.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Check wether the request method is POST.
	if r.Method != "POST" {
		// If it's not send a 405 status code in response.
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
