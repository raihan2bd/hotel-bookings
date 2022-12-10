package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/raihan2bd/hotel-go/internal/config"
	"github.com/raihan2bd/hotel-go/internal/models"
)

var app *config.AppConfig
var functions = template.FuncMap{}
var pathToTempaltes = "./templates"

// NewTemplate passing the appconfig data
func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefaultValue(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	// initilizing the template cache
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache form load case
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()

	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("fail to generate the template")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultValue(td, r)

	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal("fail to sarve the template", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// var myCache *template.Template
	myCache := map[string]*template.Template{}

	// find all pages from templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*page.html", pathToTempaltes))

	// handling err
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		// store the template file name
		name := filepath.Base(page)

		// analaize the template
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// find the layout template from templates
		matches, err := filepath.Glob(fmt.Sprintf("%s/*layout.html", pathToTempaltes))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// join the dot form layout template
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*layout.html", pathToTempaltes))
			if err != nil {
				return myCache, err
			}
		}
		// append template cache to myCache
		myCache[name] = ts
	} // end of for loop
	return myCache, nil
}
