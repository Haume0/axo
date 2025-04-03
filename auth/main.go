package auth

import (
	"axo/database"
	"axo/models"
	"fmt"
)

/*JWT Based custom authentication service
This service provides a simple way to authenticate users using JWT tokens.
*/

func Init() {
	// Initialize the authentication service
	// This function can be used to set up any necessary configurations or dependencie
	defaultRoles()
	fmt.Println("✅ Authentication service initialized")
}

func defaultRoles() {
	var dbroles []models.Role
	// Check if the roles table is empty
	if database.DB.Find(&dbroles).RowsAffected != 0 {
		return
	}
	var defaultRoles []models.Role
	//Default roles and permissions
	defaultRoles = append(defaultRoles, models.Role{
		ID:   2,
		Name: "admin",
		Permissions: []models.Permission{
			{Method: "*", Path: "*"},
		},
	})
	defaultRoles = append(defaultRoles, models.Role{
		ID:          1,
		Name:        "default",
		Permissions: []models.Permission{},
	})
	// Save the default roles to the database
	for _, role := range defaultRoles {
		// Check if role already exists
		var existingRole models.Role
		if database.DB.Where("name = ?", role.Name).First(&existingRole).RowsAffected == 0 {
			// Create new role
			database.DB.Create(&role)
		} else {
			// Update existing role's permissions
			database.DB.Model(&existingRole).Association("Permissions").Replace(role.Permissions)
		}
	}
}
