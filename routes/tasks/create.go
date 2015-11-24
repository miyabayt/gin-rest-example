package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miyabayt/gin-rest/db"
	"github.com/miyabayt/gin-rest/logger"
	"github.com/miyabayt/gin-rest/models"
)

type CreateTaskForm struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
}

func Create(c *gin.Context) {
	var form CreateTaskForm

	if err := c.Bind(&form); err == nil {
		task := &models.Task{Title: form.Title, Description: form.Description}
		db.DB.Create(task)
		logger.WithField("task", task).Infof("task created.")

		c.JSON(http.StatusOK, gin.H{"task": task})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}
