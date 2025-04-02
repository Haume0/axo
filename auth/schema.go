package auth

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Username  string    `json:"username" gorm:"type:varchar(50);unique;not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	FirstName string    `json:"first_name" gorm:"type:varchar(50)"`
	LastName  string    `json:"last_name" gorm:"type:varchar(50)"`
	Role      string    `json:"role" gorm:"type:varchar(20);default:'user'"`
	Active    bool      `json:"active" gorm:"default:true"`
}
