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

func GetAllStatusTasks(c *gin.Context) {
	var (
		result gin.H
	)

	statusTasks, err := repository.GetAllStatusTasks(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": nil,
		}
	} else {
		result = gin.H{
			"result": statusTasks,
			"count":  len(statusTasks),
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetStatusTask(c *gin.Context) {
	var (
		result gin.H
	)

	id, _ := strconv.Atoi(c.Param("id"))
	statusTask, err := repository.GetStatusTask(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": nil,
		}
	} else {
		result = gin.H{
			"result": statusTask,
		}
	}

	c.JSON(http.StatusOK, result)
}

func CreateStatusTask(c *gin.Context) {
	var statusTask structs.StatusTask

	err := c.ShouldBindJSON(&statusTask)
	if err != nil {
		panic(err)
	}

	statusTask.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	statusTask.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	err = repository.CreateStatusTask(database.DbConnection, statusTask)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status created successfully",})
}

func UpdateStatusTask(c *gin.Context) {
	var statusTask structs.StatusTask

	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&statusTask)
	if err != nil {
		panic(err)
	}

	statusTask.ID = int64(id)
	statusTask.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	err = repository.UpdateStatusTask(database.DbConnection, statusTask)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully",})
}

func DeleteStatusTask(c *gin.Context) {
	var statusTask structs.StatusTask

	id, _ := strconv.Atoi(c.Param("id"))

	statusTask.ID = int64(id)

	err := repository.DeleteStatusTask(database.DbConnection, statusTask)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status deleted successfully",})
}