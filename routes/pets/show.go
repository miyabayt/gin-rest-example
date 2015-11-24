package pets

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

	var pet models.Pet
	db.DB.First(&pet, id)
	logger.Infof("id=%d pet found.", id)

	data := map[string]interface{}{"pet": pet}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
