package models

import (
	"log"
	"taskmanager/pkg/config"
)

type Reset struct {
	UserID      int    `json:"user_id"`
	Sendcode    int    `json:"code"`
	NewPassword string `json:"newPassword"`
}

func UseResetPasswordAutomigrate() {
	db := config.GetDB()
	query := `CREATE TABLE IF NOT EXISTS reset_pass(
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    sendcode INT NOT NULL,
    newPassword VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create task table: %v", err)
	}
}

func SaveResetCode(newPass *Reset) error {
	db := config.GetDB()
	query := `INSERT INTO reset_pass(user_id,sendcode,newPassword) VALUES (?,?,?)`

	_, err := db.Exec(query, newPass.UserID, newPass.Sendcode, newPass.NewPassword)
	if err != nil {
		return err
	}

	return nil
}

func GetResetCode(id int) (int, string, error) {
	db := config.GetDB()
	query := `
		SELECT sendcode, newPassword
		FROM reset_pass
		WHERE user_id = ?
		ORDER BY id DESC
		LIMIT 1
	`

	var code int
	var password string

	err := db.QueryRow(query, id).Scan(&code, &password)
	if err != nil {
		return -1, "", err
	}
	return code, password, nil
}

func DeleteResetCode(id int) error {
	db := config.GetDB()
	query := `DELETE FROM reset_pass WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

