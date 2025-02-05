package models

import (
	"errors"

	"github.com/dalebandoni/booking-api/db"
	"github.com/dalebandoni/booking-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	q := "INSERT INTO users(email, password) VALUES (?, ?)"

	st, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer st.Close()

	hash, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	r, err := st.Exec(u.Email, hash)

	if err != nil {
		return err
	}

	userId, err := r.LastInsertId()

	u.ID = userId
	return err
}

func (u User) ValidateCreds() error {
	q := "SELECT id, password FROM users WHERE email = ?"
	r := db.DB.QueryRow(q, u.Email)

	var retreivedPassword string
	err := r.Scan(&u.ID, &retreivedPassword)

	if err != nil {
		return errors.New("Credentials are invalid.")
	}

	passwordValid := utils.CheckPasswordHash(u.Password, retreivedPassword)

	if !passwordValid {
		return errors.New("Credentials are invalid.")
	}

	return nil
}
