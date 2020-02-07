package main

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/jacobboykin/snippetbox/pkg/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// Create a human-readable date string
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

/* Create a template.FuncMap object to act as a lookup
   between the names of the custom template functions and the
   functions themeselves */
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// Parse all templates and create a template cache
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Register the template.FuncMap, then parse the file
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}