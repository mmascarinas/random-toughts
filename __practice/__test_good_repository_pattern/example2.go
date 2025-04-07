package main

import (
	"database/sql"
	"fmt"
)

type application struct {
	store Store
}

type UserRepository interface {
	GetUserByID(id int) (*User, error)
}


func bebe() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	userRepository := NewPostgresUserRepository(db)
	// userService := NewUserService(userRepository)

	// inMemRep := &InMemRepository{}

	// app := &application{
	// 	store: inMemRep,
	// }

	app := &application{
		store: userRepository,
	}

	user, err := app.store.GetUserByID(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v\n", user)
	}
}
