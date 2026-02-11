
package models

import (
	"log"
	"taskmanager/pkg/config"
)

type Task struct {
	TaskID      int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UserID      int    `json:"user_id"`
}


func TaskAutoMigrate() {
	db := config.GetDB()

	query := `CREATE TABLE IF NOT EXISTS task(
		task_id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		description TEXT,
		status VARCHAR(50) DEFAULT 'pending',
		user_id INT,
		FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create task table: %v", err)
	}
}

// Insert new task
func CreateTask(task *Task) (*Task , error) {
	db := config.GetDB()

	query := `INSERT INTO task(title, description, status, user_id) VALUES (?, ?, ?, ?);`

	_, err := db.Exec(query, task.Title, task.Description, task.Status, task.UserID)
	if err != nil {
		return nil , err
	}

	return task , nil
}

// Get task by ID
func GetTaskByID(id int) (*Task , error) {
	db := config.GetDB()

	query := `SELECT task_id, title, description, status, user_id FROM task WHERE task_id = ?;`

	task := &Task{}
	err := db.QueryRow(query, id).Scan(
		&task.TaskID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.UserID,
	)
	if err != nil {
		return nil , err
	}

	return task ,nil
}

// Get all tasks
func GetAllTasks() ([]Task, error) {
	db := config.GetDB()

	query := `SELECT task_id, title, description, status, user_id FROM task;`

	rows, err := db.Query(query)
	if err != nil {
		return nil , err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.TaskID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.UserID,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks , nil
}

// Update task by ID
func UpdateTask(task *Task) (*Task ,error) {
	db := config.GetDB()
	query := `UPDATE task SET title = ?, description = ?, status = ?, user_id = ? WHERE task_id = ?;`
	_, err := db.Exec(query, task.Title, task.Description, task.Status, task.UserID, task.TaskID)
	if err != nil {
		return nil , err
	}
	return task , nil
}

func DeleteTask(id int) error {
	db := config.GetDB()
	query := `DELETE FROM task WHERE task_id = ?;`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete task ID %d: %v", id, err)
		return err
	}
	return nil
}


