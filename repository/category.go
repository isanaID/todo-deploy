package repository

import (
	"database/sql"
	"fmt"
	"todo/structs"
)

func GetAllCategories(db *sql.DB, id int) (categories []structs.Category, err error) {
	sql := `SELECT * FROM category WHERE user_id=$1`

	fmt.Println(id)

	rows, err := db.Query(sql, id)

	if err != nil {
		return
	}

	for rows.Next() {
		var category structs.Category
		err = rows.Scan(&category.ID, &category.Name, &category.UserId, &category.CreatedAt, &category.UpdatedAt)

		if err != nil {
			return
		}

		categories = append(categories, category)
	}

	return
}

func GetCategory(db *sql.DB, id int) (category structs.Category, err error) {
	sql := `SELECT * FROM category WHERE id=$1`

	err = db.QueryRow(sql, id).Scan(&category.ID, &category.UserId, &category.Name, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return category, err
	}

	return category, nil
}

func CreateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := `INSERT INTO category (name, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4)`

	errs := db.QueryRow(sql, category.Name, category.UserId, category.CreatedAt, category.UpdatedAt)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := `UPDATE category SET name=$1, updated_at=$2 WHERE id=$3`

	errs := db.QueryRow(sql, category.Name, category.UpdatedAt, category.ID)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := `DELETE FROM category WHERE id=$1`

	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}
