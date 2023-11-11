package db

import (
	"time"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"user_name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Tickets  []Ticket
}

type Ticket struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	/* 	UserID      int */
	/* 	CreatedBy   User      `gorm:"embedded:username"` */
	DueDate time.Time `json:"due_date"`
}
