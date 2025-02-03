package models

import (
	"time"

	"github.com/dalebandoni/booking-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	q := `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES(?,?,?,?,?)`

	st, err := db.DB.Prepare(q)
	if err != nil {
		return err
	}

	defer st.Close()
	r, err := st.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := r.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	q := "SELECT * FROM events"
	rows, err := db.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	q := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(q, id)

	var e Event

	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

	if err != nil {
		return nil, err
	}

	return &e, nil
}
