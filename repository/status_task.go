package repository

import (
	"database/sql"
	"todo/structs"
)

func GetAllStatusTasks(db *sql.DB, id int) (statusTasks []structs.StatusTask, err error) {
	sql := `SELECT * FROM status_task WHERE user_id=$1`

	rows, err := db.Query(sql, id)

	if err != nil {
		return
	}

	for rows.Next() {
		var statusTask structs.StatusTask
		err = rows.Scan(&statusTask.ID, &statusTask.Status, &statusTask.UserId, &statusTask.CreatedAt, &statusTask.UpdatedAt)

		if err != nil {
			return
		}

		statusTasks = append(statusTasks, statusTask)
	}

	return
}

func GetStatusTask(db *sql.DB, id int) (statusTask structs.StatusTask, err error) {
	sql := `SELECT * FROM status_task WHERE id=$1`

	err = db.QueryRow(sql, id).Scan(&statusTask.ID, &statusTask.Status, &statusTask.CreatedAt, &statusTask.UpdatedAt)

	if err != nil {
		return statusTask, err
	}

	return statusTask, nil
}

func CreateStatusTask(db *sql.DB, statusTask structs.StatusTask) (err error) {
	sql := `INSERT INTO status_task (status, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4)`

	errs := db.QueryRow(sql, statusTask.Status, statusTask.UserId, statusTask.CreatedAt, statusTask.UpdatedAt)

	return errs.Err()
}

func UpdateStatusTask(db *sql.DB, statusTask structs.StatusTask) (err error) {
	sql := `UPDATE status_task SET status=$1, updated_at=$2 WHERE id=$3`

	errs := db.QueryRow(sql, statusTask.Status, statusTask.UpdatedAt, statusTask.ID)

	return errs.Err()
}

func DeleteStatusTask(db *sql.DB, statusTask structs.StatusTask) (err error) {
	sql := `DELETE FROM status_task WHERE id=$1`

	errs := db.QueryRow(sql, statusTask.ID)

	return errs.Err()
}