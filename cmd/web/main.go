package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/raihan2bd/bookings/internal/config"
	"github.com/raihan2bd/bookings/internal/handler"
	"github.com/raihan2bd/bookings/internal/render"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// I have to change when in production
	app.InProduction = false

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
	}

	app.TemplateCache = tc
	app.UseCache = false

	// passing the app config data
	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)

	render.NewTemplates(&app)

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
