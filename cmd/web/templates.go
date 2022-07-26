package main

import (
	"html/template"
	"path/filepath"

	"github.com/Twofold-One/tsnippet/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// newTemplateCache creates cache for templates.
func newTemplateCache() (map[string]*template.Template, error) {
	// map to act as a cache.
	cache := map[string]*template.Template{}

	// filepath.Glob() function is used to get a slice of all filepaths
	// that match the pattern ".ui/html/pages/*tmpl.html"
	pages, err := filepath.Glob("./ui/html/pages/*tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Parse the base template file into a template set.
		ts, err := template.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() *on this template set* to add any partials.
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Call ParseFiles() *on this template set* to add the page template.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map, using the name of the page as  the key.
		cache[name] = ts
	}
	return cache, nil
}
