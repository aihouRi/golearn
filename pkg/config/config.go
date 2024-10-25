package config

import "html/template"

type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
