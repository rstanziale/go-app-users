package adapters

import (
	"app-users/adapters"
	"app-users/model"
	"database/sql"
	"log"
)

var db *sql.DB

func init() {
	db = adapters.GetDb()
}

func SearchUser(email string, password string) (model.User, error) {
	log.Println("Search user")
	var user model.User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE email = $1 AND password = $2", email, password).Scan(&user.Id, &user.Name, &user.Email)
	
	return user, err
}

func SearchUserById(userId int) (model.User, error) {
	log.Println("Search user by id")
	var user model.User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", userId).Scan(&user.Id, &user.Name, &user.Email)
	
	return user, err
}

func GetUsers() ([]model.User, error) {
	log.Println("Get all users")
	var users []model.User
	rows, err := db.Query("SELECT id, name, email FROM users")
	defer rows.Close()
	
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
            log.Println(err)
    	}
		users = append(users, user)
	}

	return users, err
}

func CreateUser(name string, email string, password string) (int, error) {
	log.Println("Create user")
	var userId int
	err := db.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", name, email, password).Scan(&userId)
	
	return userId, err
}

func UpdateUser(userId int, newName string, newPassword string) error {
	log.Println("Update user")
	_, err := db.Exec("UPDATE users SET name = $1, password = $2 WHERE id = $3", newName, newPassword, userId)

	return err
}

func DeleteUser(userId int) error {
	log.Println("Delete user")
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userId)

	return err
}
