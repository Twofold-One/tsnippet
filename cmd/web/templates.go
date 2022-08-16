package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/Twofold-One/tsnippet/internal/models"
	"github.com/Twofold-One/tsnippet/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

// humanDate returns nicely formatted string representation of time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// newTemplateCache creates cache for templates.
func newTemplateCache() (map[string]*template.Template, error) {
	// map to act as a cache.
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*tmpl.html")
	if err != nil {
		return nil, err
	}

	// Code for not embeded FS
	//
	// filepath.Glob() function is used to get a slice of all filepaths
	// that match the pattern ".ui/html/pages/*tmpl.html"
	// pages, err := filepath.Glob("./ui/html/pages/*tmpl.html")
	// if err != nil {
	// 	return nil, err
	// }

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*tmpl.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		// Code for not embeded FS
		//
		// Register the template.FuncMap and parse the base template file into a template set.
		// ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		// if err != nil {
		// 	return nil, err
		// }

		// Call ParseGlob() *on this template set* to add any partials.
		// ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		// if err != nil {
		// 	return nil, err
		// }

		// Call ParseFiles() *on this template set* to add the page template.
		// ts, err = ts.ParseFiles(page)
		// if err != nil {
		// 	return nil, err
		// }

		// Add the template set to the map, using the name of the page as  the key.
		cache[name] = ts
	}

	return cache, nil
}
