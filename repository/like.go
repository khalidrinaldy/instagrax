package repository

import (
	"database/sql"
	"instagrax/structs"
)

func AddLike(db *sql.DB, like structs.Like) error {
	sql := "insert into likes (user_id, post_id) values ($1,$2)"
	err := db.QueryRow(sql, like.UserId, like.PostId)
	return err.Err()
}

func IsLiked(db *sql.DB, userId string) bool {
	var amount int
	sql := "select count(*) from likes where user_id=$1"
	row := db.QueryRow(sql, userId)
	row.Scan(&amount)
	return amount >= 1
}

func DeleteLike(db *sql.DB, id string) error {
	sql := "delete from likes where id=$1"
	err := db.QueryRow(sql, id)
	return err.Err()
}
