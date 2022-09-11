package models

import "database/sql"

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Visits int    `json:"visits"`
}

func GetUsers(db *sql.DB) (users []User, err error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return
	}

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Visits)
		if err != nil {
			break
		}

		users = append(users, user)
	}
	if err == nil {
		err = rows.Err()
	}

	return
}
