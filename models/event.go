package models

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
	UserID		int

}

var events = []Event{}

func (e Event) Save() error {

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