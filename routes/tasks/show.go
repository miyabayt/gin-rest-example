package tasks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miyabayt/gin-rest/db"
	"github.com/miyabayt/gin-rest/logger"
	"github.com/miyabayt/gin-rest/models"
)

func Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

	}

	var task models.Task
	db.DB.First(&task, id)
	logger.Infof("id=%d task found.", id)

	data := map[string]interface{}{"task": task}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
