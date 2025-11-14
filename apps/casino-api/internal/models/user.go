package models

import "time"

type User struct {
	ID        int64
	Email     string
	Password  string // hashed
	CreatedAt time.Time
}
