package repository

import (
	"database/sql"
	"instagrax/structs"
)

func AddComment(db *sql.DB, comment structs.Comment) error {
	sql := "insert into comment (text, user_id, post_id) values ($1,$2,$3)"
	err := db.QueryRow(sql, comment.Text, comment.UserId, comment.PostId)
	return err.Err()
}

func DeleteComment(db *sql.DB, id string) error {
	sql := "delete from comment where id=$1"
	err := db.QueryRow(sql, id)
	return err.Err()
}
