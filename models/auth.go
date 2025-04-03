package models

import (
	"time"
)

type User struct {
	//! Rquired fields
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
	RoleID   uint   `json:"role_id" gorm:"default:1"` // Default role is 1 (default role)
	Role     Role   `json:"role" gorm:"foreignKey:RoleID"`
	Active   bool   `json:"active" gorm:"default:true"`
	Verified bool   `json:"verified" gorm:"default:false"` // Email verified or not

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	//? Optional fields
	// Username  string `json:"username" gorm:"type:varchar(50);unique"`
	// FirstName string `json:"first_name" gorm:"type:varchar(50)"`
	// LastName  string `json:"last_name" gorm:"type:varchar(50)"`
	// Address   string `json:"address" gorm:"type:varchar(255)"`
	// Phone     string `json:"phone" gorm:"type:varchar(20)"`
	// etc. *You could create a separate table for user informations.
}

// PasswordRegex is the regex for password
var PasswordRegex string = "^[A-Za-z0-9#?!@$%^&*-.]{8,}$"

// MailRegex is the regex for mail
var MailRegex string = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

type Role struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"type:varchar(50);unique;not null"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}

type Permission struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Method string `json:"method" gorm:"type:varchar(10);not null"` // GET, POST, PUT, DELETE etc. * for all.
	Path   string `json:"path" gorm:"type:varchar(100);not null"`  // /api/v1/users etc. * for all.
}
