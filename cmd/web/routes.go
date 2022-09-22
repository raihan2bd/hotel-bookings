package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raihan2bd/bookings/internal/config"
	"github.com/raihan2bd/bookings/internal/handler"
)

func routes(app *config.AppConfig) http.Handler {
	// initializing the router form chi
	router := chi.NewRouter()

	// middleware will start here
	router.Use(SessionLoad)
	router.Use(Nosurf)

	// all routes will start here
	router.Get("/", handler.Repo.Home)
	router.Get("/about", handler.Repo.About)
	router.Get("/generals-quaters", handler.Repo.Generals)
	router.Get("/majors-suite", handler.Repo.Majors)

	router.Get("/search-availability", handler.Repo.Availability)
	router.Post("/search-availability", handler.Repo.PostAvailability)
	router.Post("/search-availability-json", handler.Repo.AvailabilityJson)

	router.Get("/contact", handler.Repo.Contact)
	router.Get("/make-reservation", handler.Repo.Reservation)

	//serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}
