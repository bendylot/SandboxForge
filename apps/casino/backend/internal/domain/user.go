package domain

import "time"

type User struct {
	ID        int64
	Login     string
	Email string
	Password  string
	CreatedAt time.Time
}
