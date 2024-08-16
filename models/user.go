package models

import (
	"log"
	"time"

	"luwa-server/utils"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Packs     []Pack    `json:"packs"`
}

func (u *User) Create() error {

	db := utils.ConnectDB()
	defer db.Close()

	err := db.QueryRow("INSERT INTO users(email, password, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING id", u.Email, u.Password, time.Now(), time.Now()).Scan(&u.ID)
	if err != nil {
		log.Panic("error creating a new user", err)
		return err
	}

	return nil
}

func (u *User) GetUserByEmail(email string) (*User, error) {

	db := utils.ConnectDB()
	defer db.Close()

	err := db.QueryRow("SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1", email).Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		log.Panic("error getting user by email", err)
		return nil, err
	}

	return u, nil
}
