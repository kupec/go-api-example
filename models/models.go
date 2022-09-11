package models

import "github.com/jmoiron/sqlx"

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Visits int    `json:"visits"`
}

func GetUsers(db *sqlx.DB) (users []User, err error) {
	err = db.Select(&users, "SELECT * FROM users")
	return
}
