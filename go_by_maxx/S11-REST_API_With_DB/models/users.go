package models

import (
	"errors"

	"example.com/rest_api_with_db/db"
	"example.com/rest_api_with_db/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (U *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(&U.Password)

	if err != nil {
		return err
	}

	U.Password = hashedPassword

	result, err := stmt.Exec(&U.Email, &U.Password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	U.Id = userId

	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id ,&retrievedPassword)

	if err != nil {
		return errors.New("Credentials are invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials are invalid")
	}

	return nil
}

// UserExists checks if a user with the given ID exists in the database
func UserExists(userId int64) (bool, error) {
	query := `SELECT id FROM users WHERE id = ?`
	row := db.DB.QueryRow(query, userId)

	var id int64
	err := row.Scan(&id)

	if err != nil {
		return false, err
	}

	return true, nil
}
