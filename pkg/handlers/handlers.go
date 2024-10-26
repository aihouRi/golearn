package handlers

import (
	"net/http"

	"github.com/aihouRi/golearn/pkg/config"
	"github.com/aihouRi/golearn/pkg/models"
	"github.com/aihouRi/golearn/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.Appconfig
}

// NewRepo create a new repository
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is a home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is a about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "Hello, again."
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: StringMap,
	})
}
