package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	InProduction  bool
	UseCache      bool
	InfoLog       *log.Logger
	Session       *scs.SessionManager
}
