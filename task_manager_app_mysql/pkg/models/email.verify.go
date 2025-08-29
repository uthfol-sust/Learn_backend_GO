package models

import (
	"fmt"
	"taskmanager/pkg/config"
	"time"
	"log"
)

type EmailVerify struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`   
	Email     string    `json:"email"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
	IsVerified bool     `json:"is_verified"`
}

func EmailVerificationAutoMigrate(){
	db := config.GetDB()
    
	query := `CREATE TABLE IF NOT EXISTS email_verifications(
	         id INT AUTO_INCREMENT PRIMARY KEY,
			 user_id INT NOT NULL,
			 email VARCHAR(100) NOT NULL,
			 code VARCHAR(10) NOT NULL,
			 expires_at TIMESTAMP NOT NULL,
             is_verified BOOLEAN DEFAULT FALSE,
			  FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE
	       );`


	_ , err :=db.Exec(query)

	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
}

func SaveVerification(userId int, Email , Code string, ExpiresAt time.Time) error {
    db := config.GetDB()

    query :=`INSERT INTO email_verifications (user_id, email, code ,expires_at) VALUES (?, ?, ?, ?)`

	_ , err := db.Exec(query, userId, Email, Code, ExpiresAt)
	return err
}

func GetVerificationByEmail(email string) (*EmailVerify, error) {
	db := config.GetDB()

	query := `SELECT  code, expires_at , is_verified
	          FROM email_verifications WHERE email = ? 
	          ORDER BY id DESC LIMIT 1`

	row := db.QueryRow(query, email)

	var expiresAtStr string

	ev := &EmailVerify{}
	err := row.Scan( &ev.Code, &expiresAtStr, &ev.IsVerified)
	if err != nil {
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	ev.ExpiresAt, _ = time.Parse("2006-01-02 15:04:05", expiresAtStr)

	return ev, nil
}



func MarkVerified(email string) error {
	db := config.GetDB()
	query := `UPDATE email_verifications 
	          SET is_verified = TRUE 
	          WHERE email = ? 
	          ORDER BY id DESC LIMIT 1`

	res, err := db.Exec(query, email)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no verification record found for %s", email)
	}

	return nil
}
