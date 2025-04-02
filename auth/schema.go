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
	RoleID    uint      `json:"role_id"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleID"`
	Active    bool      `json:"active" gorm:"default:true"`
}

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
