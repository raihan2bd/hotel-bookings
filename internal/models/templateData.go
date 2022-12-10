package models

import "github.com/raihan2bd/hotel-go/internal/forms"

// TemplateData sent data from handler to templates

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Warning   string
	Flash     string
	Error     string
	Form      *forms.Form
}
