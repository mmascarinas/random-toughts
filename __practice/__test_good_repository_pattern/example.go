package main

import (
	"database/sql"
	"fmt"
)

type Store interface {
	GetUserByID(id int) (*User, error)
}

type InMemRepository struct {
	users []User
}

func (r *InMemRepository) GetUserByID(id int) (*User, error) {
	return nil, nil
}

type User struct {
	ID   int
	Name string
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetUserByID(id int) (*User, error) {
	row := r.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

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

	userRepository := NewPostgresUserRepository(db)
	// userService := NewUserService(userRepository)

	user, err := userRepository.GetUserByID(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v\n", user)
	}
}


