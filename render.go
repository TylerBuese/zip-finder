package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, t string) {
	cache, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	template, ok := cache[t]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = template.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.html.template")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)

		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return cache, err
			}
		}

		cache[name] = templateSet
	}

	return cache, nil

}
