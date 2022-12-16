package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todo/database"
	"todo/repository"
	"todo/structs"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	var (
		result gin.H
	)

	id, _ := strconv.Atoi(c.Query("user_id"))
	fmt.Println(id)

	tasks, err := repository.GetAllTasks(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": nil,
		}
	} else {
		result = gin.H{
			"result": tasks,
			"count":  len(tasks),
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetTask(c *gin.Context) {
	var (
		result gin.H
	)

	id, _ := strconv.Atoi(c.Param("id"))
	task, err := repository.GetTask(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": nil,
		}
	} else {
		result = gin.H{
			"result": task,
		}
	}

	c.JSON(http.StatusOK, result)
}

func CreateTask(c *gin.Context) {
	var task structs.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		panic(err)
	}

	task.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	err = repository.CreateTask(database.DbConnection, task)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task created successfully",
	})
}

func UpdateTask(c *gin.Context) {
	var task structs.Task

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&task)
	if err != nil {
		panic(err)
	}

	task.ID = int64(id)
	task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	err = repository.UpdateTask(database.DbConnection, task)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
	})
}

func DeleteTask(c *gin.Context) {
	var task structs.Task

	id, _ := strconv.Atoi(c.Param("id"))

	task.ID = int64(id)

	err := repository.DeleteTask(database.DbConnection, task)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}