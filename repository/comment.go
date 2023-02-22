package repository

import (
	"database/sql"
	"instagrax/structs"
)

func GetAllComment(db *sql.DB, id string) (comments []structs.Comment, err error) {
	sql := "select * from comment where post_id=$1"
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment structs.Comment
		err = rows.Scan(&comment.Id, &comment.Text, &comment.PostId, &comment.UserId, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			panic(err)
		}
		comments = append(comments, comment)
	}
	return
}

func AddComment(db *sql.DB, comment structs.Comment) error {
	sql := "insert into comments (text, user_id, post_id) values ($1,$2,$3)"
	err := db.QueryRow(sql, comment.Text, comment.UserId, comment.PostId)
	return err.Err()
}

func DeleteComment(db *sql.DB, id string) error {
	sql := "delete from comments where id=$1"
	err := db.QueryRow(sql, id)
	return err.Err()
}
