package repository

import (
	"database/sql"
	"instagrax/structs"
)

func AddLike(db *sql.DB, like structs.Like) error {
	sql := "insert into like (user_id, post_id) values ($1,$2)"
	err := db.QueryRow(sql, like.UserId, like.PostId)
	return err.Err()
}

func DeleteLike(db *sql.DB, id string) error {
	sql := "delete from like where id=$1"
	err := db.QueryRow(sql, id)
	return err.Err()
}
