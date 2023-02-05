package repository

import "github.com/raihan2bd/hotel-go/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}
