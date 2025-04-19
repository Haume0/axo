package auth

import (
	"axo/database"
	"axo/models"
	"slices"
)

// Role operations
func CreateRole(name string, permissions []models.Permission) error {
	var role models.Role = models.Role{
		Name:        name,
		Permissions: permissions,
	}
	if err := database.DB.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRole(roleID uint) error {
	if err := database.DB.Delete(&models.Role{}, roleID).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRole(roleID uint, name string, permissions []models.Permission) error {
	var role models.Role
	if err := database.DB.First(&role, roleID).Error; err != nil {
		return err
	}
	role.Name = name
	role.Permissions = permissions
	if err := database.DB.Save(&role).Error; err != nil {
		return err
	}
	return nil
}

func GetRole(roleID uint) (models.Role, error) {
	var role models.Role
	if err := database.DB.First(&role, roleID).Error; err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func ListRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := database.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// Permission operations
func AddPermissionToRole(roleID uint, permission models.Permission) error {
	var role models.Role
	if err := database.DB.First(&role, roleID).Error; err != nil {
		return err
	}
	role.Permissions = append(role.Permissions, permission)
	if err := database.DB.Save(&role).Error; err != nil {
		return err
	}
	return nil
}

func RemovePermissionFromRole(roleID uint, permissionID uint) error {
	var role models.Role
	if err := database.DB.First(&role, roleID).Error; err != nil {
		return err
	}
	for i, perm := range role.Permissions {
		if perm.ID == permissionID {
			role.Permissions = slices.Delete(role.Permissions, i, i+1)
			break
		}
	}
	if err := database.DB.Save(&role).Error; err != nil {
		return err
	}
	return nil
}

func CheckPermission(userID uint, method string, path string) (bool, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return false, err
	}
	var role models.Role
	if err := database.DB.First(&role, user.RoleID).Error; err != nil {
		return false, err
	}
	for _, permission := range role.Permissions {
		if permission.Method == method && permission.Path == path {
			return true, nil
		}
	}
	return false, nil
}
