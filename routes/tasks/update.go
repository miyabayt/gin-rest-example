package tasks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miyabayt/gin-rest/db"
	"github.com/miyabayt/gin-rest/logger"
	"github.com/miyabayt/gin-rest/models"
)

type UpdateTaskForm struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}

func Update(c *gin.Context) {
	var form UpdateTaskForm

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

	}

	if err := c.Bind(&form); err == nil {
		task := &models.Task{Title: form.Title, Description: form.Description}
		db.DB.First(task, id).Update(task)
		logger.WithField("task", task).Infof("task saved.")

		c.JSON(http.StatusOK, gin.H{"task": task})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}
