package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Twofold-One/tsnippet/internal/models"
)

// home handler function which writes a byte slice containing
// "Hello from tsnippet" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if current request URL path exactly mathces "/".
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "home.tmpl.html", &templateData{
		Snippets: snippets,
	})
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
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, http.StatusOK, "view.tmpl.html", &templateData{
		Snippet: snippet,
	})
}

// snippetCreate handler function which creates a new snippet.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Check wether the request method is POST.
	if r.Method != http.MethodPost {
		// If it's not send a 405 status code in response.
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	// dummy data for now
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\nKobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
