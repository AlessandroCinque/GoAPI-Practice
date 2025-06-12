package modelsWithDBQueries

import (
	"time"

	"github.com/AlessandroCinque/GoAPI-Practice/db"
)

type Event struct {
	ID          int64
	Name        string		`binding:"required"`
	Description string		`binding:"required"`
	Location    string		`binding:"required"`
	DateTime    time.Time	`binding:"required"`
	UserID		int64

}

var events = []Event{}

//The pointer "*" in here is required, so that we update the actual user not a copy of it
func (e *Event) Save() error {

	query := `INSERT INTO events(name,description,location, dateTime, user_id)
	VALUES(?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)
	
	
	if err != nil {
		return err
	}
	
	defer stmt.Close()

	//Exec is commonly used to modify/insert data
	
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {

	//Query is commonly used to retrieve data
	query := "SELECT * FROM events"
	rows, err :=db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	var events []Event

	//Next returns true as long as there are still rows
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name,&event.Description,&event.Location, &event.DateTime, &event.UserID)
		
		if err != nil {
			return nil,err
		}
		events = append(events, event)

	}

	
	defer rows.Close()
	return events, nil
}
//Event itself can't be null but a POINTER to an event can
func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	//God for 1 to 1
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name,&event.Description,&event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil ,err
	}

	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET	name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err :=db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	stmt.Exec(e.ID, e.UserID)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	stmt.Exec(e.ID, e.UserID)

	return err
}