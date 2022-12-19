package repository

import (
	"database/sql"
	"todo/structs"
)

func RegisterUser(db *sql.DB, user structs.User) (err error) {
	sql := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	errs := db.QueryRow(sql, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)

	return errs.Err()
}

func CheckEmail(db *sql.DB, email string) (err error) {
	sql := `SELECT email FROM users WHERE email=$1`

	err = db.QueryRow(sql, email).Scan(&email)

	if err != nil {
		return err
	}

	return nil
}

func LoginUser(db *sql.DB, email string) (user structs.User, err error) {
	sql := `SELECT * FROM users WHERE email=$1`

	err = db.QueryRow(sql, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}