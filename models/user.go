package models

import (
	"errors"

	"github.com/AlessandroCinque/GoAPI-Practice/db"
	"github.com/AlessandroCinque/GoAPI-Practice/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	
	hashedPass, err := utils.HashPasssword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPass)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}


//The pointer "*" in here is required, so that we update the actual user not a copy of it
func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievePassword string
	err := row.Scan(&u.ID,&retrievePassword)

	if err != nil {
		return err
	}

	isPassValid := utils.CHeckPassHash(u.Password,retrievePassword)

	if !isPassValid {
		return errors.New("Credentials invalid")
	}

	return nil
}