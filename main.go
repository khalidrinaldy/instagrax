package main

import (
	"database/sql"
	"fmt"
	"instagrax/database"
	"instagrax/route"
	"os"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))
	//os.Getenv("DB_HOST"),
	//os.Getenv("DB_PORT"),
	//os.Getenv("DB_USER"),
	//os.Getenv("DB_PASSWORD"),
	//os.Getenv("DB_DATABASE"))
	fmt.Println(psqlInfo)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Failed connect to database")
		panic(err)
	} else {
		fmt.Println("Success connect to database")
	}
	defer DB.Close()

	database.DbMigrate(DB)
	route.StartServer().Run(":" + os.Getenv("PORT"))
}
