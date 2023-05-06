package handlers

import (
	"net/http"

	"github.com/simonwardle/bookings/pkg/config"
	"github.com/simonwardle/bookings/pkg/models"
	"github.com/simonwardle/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepository creates a new repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (rec *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rec.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.tmpl.html", &models.TemplateData{})
}

// About is the about page handler
func (rec *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again!"

	remoteIP := rec.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP 

	render.RenderTemplate(w, "about.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals is the generals page handler
func (rec *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.tmpl.html", &models.TemplateData{
	})
}

// Majors is the majorss page handler
func (rec *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.tmpl.html", &models.TemplateData{
	})
}

// Availability search for room availability
func (rec *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.tmpl.html", &models.TemplateData{
	})
}

// Availability search for room availability
func (rec *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.tmpl.html", &models.TemplateData{
	})
}

// Reservation capture client details
func (rec *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.tmpl.html", &models.TemplateData{
	})
}