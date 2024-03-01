package adapters

import (
	"app-users/adapters"
	"app-users/model"
	"log"
)

type UserAdapter struct {
	dao adapters.DAO
}

// UserAdapter constructor
func NewUserAdapter() *UserAdapter {
	adapter := UserAdapter{adapters.GetDAO()}
	return &adapter
}

// Serach user by email and password
func (adapter *UserAdapter) SearchUser(email string, password string) (model.User, error) {
	log.Println("Search user")

	var user model.User
	err := adapter.dao.GetDb().QueryRow("SELECT id, name, email FROM users WHERE email = $1 AND password = $2", email, password).Scan(&user.Id, &user.Name, &user.Email)

	return user, err
}

// Serach user by its identifier
func (adapter *UserAdapter) SearchUserById(userId int) (model.User, error) {
	log.Println("Search user by id")

	var user model.User
	err := adapter.dao.GetDb().QueryRow("SELECT id, name, email FROM users WHERE id = $1", userId).Scan(&user.Id, &user.Name, &user.Email)

	return user, err
}

// Get all registred users
func (adapter *UserAdapter) GetUsers() ([]model.User, error) {
	log.Println("Get all users")

	var users []model.User
	rows, err := adapter.dao.GetDb().Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}

// Create new user
func (adapter *UserAdapter) CreateUser(name string, email string, password string) (int, error) {
	log.Println("Create user")

	var userId int
	err := adapter.dao.GetDb().QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", name, email, password).Scan(&userId)

	return userId, err
}

// Update user info
func (adapter *UserAdapter) UpdateUser(userId int, newName string, newPassword string) error {
	log.Println("Update user")

	_, err := adapter.dao.GetDb().Exec("UPDATE users SET name = $1, password = $2 WHERE id = $3", newName, newPassword, userId)

	return err
}

// Delete a user
func (adapter *UserAdapter) DeleteUser(userId int) error {
	log.Println("Delete user")

	_, err := adapter.dao.GetDb().Exec("DELETE FROM users WHERE id = $1", userId)

	return err
}
