package auth

import "axo/models"

// Role operations
func CreateRole(name string, permissions []models.Permission) error {
	// TODO: Will create a new role with the given permissions.
	return nil
}
func DeleteRole(roleID uint) error {
	// TODO: Will delete the role by ID.
	return nil
}
func UpdateRole(roleID uint, name string, permissions []models.Permission) error {
	// TODO: Will update role name and permissions.
	return nil
}
func GetRole(roleID uint) (models.Role, error) {
	// TODO: Will return role details by ID.
	return models.Role{}, nil
}
func ListRoles() ([]models.Role, error) {
	// TODO: Will return all available roles.
	return []models.Role{}, nil
}

// Permission operations
func AddPermissionToRole(roleID uint, permission models.Permission) error {
	// TODO: Will add a specific permission to the given role.
	return nil
}
func RemovePermissionFromRole(roleID uint, permissionID uint) error {
	// TODO: Will remove a specific permission from the given role.
	return nil
}
func CheckPermission(userID uint, method string, path string) (bool, error) {
	// TODO: Will check if the user has permission for the requested method and path.
	return false, nil
}
