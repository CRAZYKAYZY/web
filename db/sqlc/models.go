// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
)

type User struct {
	ID             int32        `json:"id"`
	Username       string       `json:"username"`
	HashedPassword string       `json:"hashed_password"`
	FullName       string       `json:"full_name"`
	Email          string       `json:"email"`
	CreatedAt      sql.NullTime `json:"created_at"`
}
