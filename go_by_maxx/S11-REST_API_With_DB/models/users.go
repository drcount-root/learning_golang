package models

import (
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

func GetUser(email, password *string) (*User,error) {
	query := `SELECT * FROM users WHERE email = ? AND password = ?`
	row := db.DB.QueryRow(query, email, password)

	var user User
	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
