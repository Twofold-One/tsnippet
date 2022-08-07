package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Twofold-One/tsnippet/internal/models"
	"github.com/Twofold-One/tsnippet/internal/validator"
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
	data := app.newTemplateData(r)

	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, http.StatusOK, "create.tmpl.html", data)
}

type snippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
}

// snippetCreatePost handler function which creates a new snippet.
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	var form snippetCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.formDecoder.Decode(&form, r.PostForm)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.ChekField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.ChekField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.ChekField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.ChekField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must be equal 1, 7, or 365")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display a HTML form for signing up a new user...")
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user...")
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display a HTML form for logging in a new user...")
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
