package repository

import (
	"database/sql"
	"errors"
	"instagrax/structs"
)

func GetAllUsers(db *sql.DB) (users []structs.User, err error) {
	sql := "select * from user"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user structs.User
		err := rows.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return
}

func CheckEmail(db *sql.DB, email string) (user structs.User, err error) {
	var users []structs.User
	sql := "select * from user where email=$1"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	if len(users) == 1 {
		return
	} else {
		err = errors.New("email belum terdaftar")
		return
	}
}

func Register(db *sql.DB, user structs.User) error {
	sql := "insert into user (username, name, email, password) values ($1,$2,$3,$4)"
	errs := db.QueryRow(sql, user.Username, user.Name, user.Email, user.Password)
	return errs.Err()
}
