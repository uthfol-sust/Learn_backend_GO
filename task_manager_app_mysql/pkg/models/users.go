package models

import (
	"log"
	"database/sql"
	"taskmanager/pkg/config"
)

// User model
type User struct {
	UserID   int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// Auto-migrate users table
func UserAutoMigrate() {
	db := config.GetDB()

	query := `CREATE TABLE IF NOT EXISTS user(
		user_id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		role VARCHAR(50) DEFAULT 'user'
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
}

// Register a new user
func UserRegistration(user *User) (*User, error){
	db := config.GetDB()

	query := `INSERT INTO user (name, password, email, role) VALUES (?, ?, ?, ?);`

	result, err:= db.Exec(query, user.Name, user.Password, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	user.UserID = int(id)

	return user, nil
}

// Get user by ID
func GetUserByID(id int) (*User , error ){
	db := config.GetDB()

	query := `SELECT user_id, name, email, password, role FROM user WHERE user_id = ?;`

	user := &User{}
	err := db.QueryRow(query, id).Scan(
		&user.UserID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return nil, err
	}

	return user ,nil
}

// Get all users
func GetAllUsers() ( []User , error ){
	db := config.GetDB()

	query := `SELECT user_id, name, email, role FROM user`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.UserID,
			&user.Name,
			&user.Email,
			&user.Role,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil , err
	}

	return users ,nil
}


func UpdateUser(user *User) error {
	db := config.GetDB()
    
	query := `UPDATE user SET name=?,password=? WHERE user_id=?`

	_, err := db.Exec(query, user.Name , user.Password,user.UserID)

	return err
}

func DeleteUser(ID int) error {
	db := config.GetDB()

	query := `DELETE FROM user WHERE user_id=?`

	result, err := db.Exec(query, ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

