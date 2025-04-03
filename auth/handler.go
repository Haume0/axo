package auth

import (
	"axo/axo"
	"axo/database"
	"axo/models"
	"crypto/sha256"
	"fmt"
)

// Authentication operations
func Logout() {
	// TODO: Will logout the user.
}

func Login() {
	// TODO: Will login the user.
}

func Register(user models.User) error {
	//Check user.Mail with MailRegex
	if !axo.Unwrap(axo.RegexTest(user.Email, models.MailRegex)) {
		return fmt.Errorf("BAD_MAIL_FORMAT")
	}

	//Check user.Password with PasswordRegex
	if !axo.Unwrap(axo.RegexTest(user.Password, models.PasswordRegex)) {
		return fmt.Errorf("BAD_PASSWORD_FORMAT")
	}

	//Checking if user exists,
	var searchUsers []models.User
	database.DB.Where("email = ?", user.Email).Find(&searchUsers)
	if len(searchUsers) > 0 {
		return fmt.Errorf("USER_ALREADY_EXISTS")
	}

	//sha256 hash the password
	sha := sha256.New()
	sha.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", sha.Sum(nil))

	//Creating user
	database.DB.Create(&user)
	return nil
}

func Refresh() {
	// TODO: Will refresh the user token.
}

func Verify() {
	// TODO: Will verify the user email.
}

// User operations
func CheckAuth() (bool, error) {
	// TODO: Will check if user is authenticated or not.
	return false, nil
}
func GetUser() (models.User, error) {
	// TODO: Will return user data if authenticated.
	return models.User{}, nil
}

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
