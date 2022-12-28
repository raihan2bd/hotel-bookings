package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/raihan2bd/hotel-go/internal/config"
	"github.com/raihan2bd/hotel-go/internal/handler"
	"github.com/raihan2bd/hotel-go/internal/helpers"
	"github.com/raihan2bd/hotel-go/internal/models"
	"github.com/raihan2bd/hotel-go/internal/render"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	// initilizing the server
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	fmt.Printf("Server is running on http://localhost%s \n", port)

	err = srv.ListenAndServe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Fatal(err)

}

func run() error {
	// What am i going to put the session
	gob.Register(models.Reservation{})

	// I have to change when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cann't create template cash")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	// passing the app config data
	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)
	return nil
}
