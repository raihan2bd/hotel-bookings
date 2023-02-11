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
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	var newID int
	stmt := `INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return newID, nil
}

// InsertRoomRestrictions is insert room restriction
func (m *postgresDBRepo) InsertRoomRestrictions(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	stmt := `insert into room_restrictions (start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		r.RestrictionID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}

// SearchAvailabilityByDateByRoomID return true if availability exist for roomID, and false if no availablity
func (m *postgresDBRepo) SearchAvailabilityByDateByRoomID(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	query := `
	select
		count(id)
	from
		room_restrictions
	where
		room_id = $1
		'$2 <end_date and $3 > start_date;`

	var numRows int
	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}

	return false, nil
}

// SearchAvailabilityForAllRooms return a slice of all avilable rooms
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `
	select
		r.id, r.room_name
	from
		rooms r
	where r.id not in
		(select room_id from room_restrictions rr where $1 < rr.end_date and $2 > rr.start_date);
	`

	rows, err := m.DB.QueryContext(ctx, query, start, end)

	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)

		if err != nil {
			return rooms, err
		}

		rooms = append(rooms, room)

	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}
