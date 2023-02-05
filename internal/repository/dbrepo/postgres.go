package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/raihan2bd/hotel-go/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// insert reservation inside reservation
func (m *postgresDBRepo) InsertReservation(res models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	fmt.Println(res.RoomID)

	stmt := `INSERT INTO reservaions (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.ExecContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
