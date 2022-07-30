package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Twofold-One/tsnippet/internal/models"
	"github.com/julienschmidt/httprouter"
)

// home handler function which writes a byte slice containing
// "Hello from tsnippet" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.tmpl.html", data)
}

// snippetView handler function which displays a specific snippet.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	// Extract the value of "id" parameter from the params slice
	// convert it to an integer and check for errors.
	id, err := strconv.Atoi(params.ByName("id"))
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

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.tmpl.html", data)
}

// snippetCreate handler shows a HTML form to create a new snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the form for creating a new snippet..."))
}

// snippetCreatePost handler function which creates a new snippet.
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// dummy data for now
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\nKobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
