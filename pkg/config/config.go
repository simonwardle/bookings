package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application configuation
type AppConfig struct {
	UseTemplateCache bool
	TemplateCache    map[string]*template.Template
	InfoLog          *log.Logger
	InProduction     bool
	Session          *scs.SessionManager
}
