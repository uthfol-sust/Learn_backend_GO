package models

import (
	"database/sql"
	"fmt"
	"log"
	"taskmanager/pkg/config"
)


type User struct {
	UserID   int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

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


func GetUserByID(id int) (*User , error ){
	db := config.GetDB()

	query := `SELECT user_id, name, password, email, role FROM user WHERE user_id = ?;`

	user := &User{}
	err := db.QueryRow(query, id).Scan(&user.UserID, &user.Name,&user.Password, &user.Email, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with id:", id)
			return nil, nil
		}
		return nil, err
	}

	return user ,nil
}


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
    
	query := `UPDATE user SET name=?,password=?, role=? WHERE user_id=?`

	_, err := db.Exec(query, user.Name , user.Password, user.Role, user.UserID)

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


func FindUserByEmail(email string)(*User , error){
	db := config.GetDB()

	query := `SELECT user_id, password, email, role FROM user WHERE email=?;`

	user := &User{}
	err := db.QueryRow(query,email).Scan(&user.UserID,&user.Password, &user.Email ,&user.Role)

	return user , err
}

