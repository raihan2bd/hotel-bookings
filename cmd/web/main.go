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
	"github.com/raihan2bd/hotel-go/internal/driver"
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
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	db.SQL.Close()

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

func run() (*driver.DB, error) {
	// What am i going to put the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

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

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=booking user=postgres password=password")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	log.Println("connected to the database")

	// template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cann't create template cash")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	// passing the app config data
	repo := handler.NewRepo(&app, db)
	handler.NewHandler(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)
	return db, nil
}
