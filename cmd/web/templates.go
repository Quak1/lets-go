package main

import (
	"html/template"
	"path/filepath"

	"github.com/Quak1/snippetbox/internal/models"
)

type templateData struct {
	Snippet     models.Snippet
	Snippets    []models.Snippet
	CurrentYear int
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		filename := filepath.Base(page)

		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(page)
		if err != nil {
			return nil, err
		}

		cache[filename] = ts
	}

	return cache, nil
}
