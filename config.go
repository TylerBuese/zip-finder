package main

import "text/template"

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
