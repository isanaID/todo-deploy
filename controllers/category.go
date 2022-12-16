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

	categories, err := repository.GetAllCategories(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": nil,
		}
	} else {
		result = gin.H{
			"result": categories,
			"count":  len(categories),
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetCategory(c *gin.Context) {
	var (
		result gin.H
	)

	id, _ := strconv.Atoi(c.Param("id"))
	category, err := repository.GetCategory(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": nil,
		}
	} else {
		result = gin.H{
			"result": category,
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

	category.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	category.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	err = repository.CreateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category created successfully",
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
		"message": "Category deleted successfully",
	})
}
