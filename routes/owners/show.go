package owners

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

	var owner models.Owner
	db.DB.First(&owner)
	logger.Infof("id=%d owner found.", id)

	data := map[string]interface{}{"owner": owner}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
