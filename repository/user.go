package repository

import (
	"database/sql"
	"errors"
	"instagrax/structs"
)

func GetAllUsers(db *sql.DB) (users []structs.User, err error) {
	sql := "select * from users"
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
	sql := "select * from users where email=$1"
	rows, err := db.Query(sql, email)
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

func CheckUsername(db *sql.DB, username string) (user structs.User, err error) {
	var users []structs.User
	sql := "select * from users where username=$1"
	rows, err := db.Query(sql, username)
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
		err = errors.New("username belum terdaftar")
		return
	}
}

func CheckId(db *sql.DB, id string) (user structs.User, err error) {
	var users []structs.User
	sql := "select * from users where id=$1"
	rows, err := db.Query(sql, id)
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
		err = errors.New("id tidak terdaftar")
		return
	}
}

func Register(db *sql.DB, user structs.User) error {
	sql := "insert into users (username, name, email, password) values ($1,$2,$3,$4)"
	errs := db.QueryRow(sql, user.Username, user.Name, user.Email, user.Password)
	return errs.Err()
}

func EditProfile(db *sql.DB, user structs.User) error {
	sql := "update users set username=$1, name=$2 where id=$3"
	errs := db.QueryRow(sql, user.Username, user.Name, user.Id)
	return errs.Err()
}
