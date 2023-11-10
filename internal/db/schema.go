package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"user_name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Tickets  []Ticket
}

type Ticket struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int
	CreatedBy   User      `gorm:"embedded:username"`
	DueDate     time.Time `json:"due_date"`
}
