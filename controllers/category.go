package controllers

import (
	"net/http"
	"strconv"
	"time"
	"todo/database"
	"todo/repository"
	"todo/structs"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	var (
		result gin.H
	)
	id := IdUser(c)

	categories, err := repository.GetAllCategories(database.DbConnection, id)

	type response struct {
		Name string `json:"name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	if err != nil {
		result = gin.H{
			"result": nil,
		}
	} else {
		result = gin.H{
			"data": categories,
			"count":  len(categories),
		}
	}

	c.JSON(http.StatusOK, result)
}

func CreateCategory(c *gin.Context) {
	var category structs.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	if category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "Category name is required",
		})
		return
	}

	id := IdUser(c)

	category.UserId = int64(id)
	category.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	category.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	err = repository.CreateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Category Created Successfully",
	})
}

func UpdateCategory(c *gin.Context) {
	var category structs.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}
	id, _ := strconv.Atoi(c.Param("id"))
	category.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	category.ID = int64(id)

	err = repository.UpdateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Category updated successfully",
	})
}

func DeleteCategory(c *gin.Context) {
	var category structs.Category

	id, _ := strconv.Atoi(c.Param("id"))
	category.ID = int64(id)

	err := repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Category deleted successfully",
	})
}
