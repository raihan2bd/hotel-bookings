package handler

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"html/template"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/justinas/nosurf"
	"github.com/raihan2bd/hotel-go/internal/config"
	"github.com/raihan2bd/hotel-go/internal/models"
	"github.com/raihan2bd/hotel-go/internal/render"
)

var app config.AppConfig
var session *scs.SessionManager
var functions = template.FuncMap{}
var pathToTempaltes = "./../../templates"

func getRoutes() http.Handler {
	// What am i going to put the session
	gob.Register(models.Reservation{})

	// I have to change when in production
	app.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// template cache
	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cann't create template cash")
	}

	app.TemplateCache = tc
	app.UseCache = true

	// passing the app config data
	repo := NewRepo(&app)
	NewHandler(repo)

	render.NewTemplates(&app)

	// initializing the router form chi
	router := chi.NewRouter()

	// middleware will start here
	router.Use(SessionLoad)
	//router.Use(Nosurf)

	// all routes will start here
	router.Get("/", Repo.Home)
	router.Get("/about", Repo.About)
	router.Get("/generals-quaters", Repo.Generals)
	router.Get("/majors-suite", Repo.Majors)

	router.Get("/search-availability", Repo.Availability)
	router.Post("/search-availability", Repo.PostAvailability)
	router.Post("/search-availability-json", Repo.AvailabilityJson)

	router.Get("/contact", Repo.Contact)
	router.Get("/make-reservation", Repo.Reservation)
	router.Post("/make-reservation", Repo.PostReservation)
	router.Get("/reservation-summary", Repo.ReservationSummary)

	//serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}

// Nosurf CSRF protection to all POST request
func Nosurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {
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
