package models

import (
	"errors"

	"github.com/howters/golang/db"
	"github.com/howters/golang/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password)
	VALUES(?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashed,err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashed)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u User) VerifyCredentials() error {
	query := `SELECT id, password FROM users
	WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}
	
	isValid := utils.CheckPasswordHash(u.Email, retrievedPassword)

	if !isValid {
		return errors.New("Invalid Credentials!")
	}

	return nil
}