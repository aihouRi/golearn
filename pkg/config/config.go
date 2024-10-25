package config

import "html/template"

type Appconfig struct {
	TemplateCache map[string]*template.Template
}