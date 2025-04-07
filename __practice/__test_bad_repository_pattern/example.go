package main

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func getUserByID(db *sql.DB, id int) (*User, error) {
	row := db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	var user User
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}

	return &user, nil
}

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	user, err := getUserByID(db, 1)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v\n", user)
	}
}
