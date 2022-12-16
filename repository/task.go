package repository

import (
	"database/sql"
	"todo/structs"
)

func GetAllTasks(db *sql.DB, id int) (tasks []structs.Task, err error) {
	sql := `SELECT * FROM task WHERE user_id=$1`

	rows, err := db.Query(sql, id)

	if err != nil {
		return 
	}

	for rows.Next() {
		var task structs.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Deadline, &task.User_id, &task.Category_id, &task.Status_id, &task.CreatedAt, &task.UpdatedAt)

		if err != nil {
			return 
		}

		tasks = append(tasks, task)
	}

	return
}

func GetTask(db *sql.DB, id int) (task structs.Task, err error) {
	sql := `SELECT * FROM task WHERE id=$1`

	err = db.QueryRow(sql, id).Scan(&task.ID, &task.Title, &task.Description, &task.Deadline, &task.User_id, &task.Category_id, &task.Status_id, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return task, err
	}

	return task, nil
}

func CreateTask(db *sql.DB, task structs.Task) (err error) {
	sql := `INSERT INTO task (title, description, deadline, user_id, category_id, status_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	errs := db.QueryRow(sql, task.Title, task.Description, task.Deadline, task.User_id, task.Category_id, task.Status_id, task.CreatedAt, task.UpdatedAt)

	return errs.Err()
}

func UpdateTask(db *sql.DB, task structs.Task) (err error) {
	sql := `UPDATE task SET title=$1, description=$2, user_id=$3, category_id=$4, status_id=$5, updated_at=$6 WHERE id=$7`

	errs := db.QueryRow(sql, task.Title, task.Description, task.User_id, task.Category_id, task.Status_id, task.UpdatedAt, task.ID)

	return errs.Err()
}

func DeleteTask(db *sql.DB, task structs.Task) (err error) {
	sql := `DELETE FROM task WHERE id=$1`

	errs := db.QueryRow(sql, task.ID)

	return errs.Err()
}