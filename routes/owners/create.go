package owners

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miyabayt/gin-rest/db"
	"github.com/miyabayt/gin-rest/logger"
	"github.com/miyabayt/gin-rest/models"
)

type OwnerForm struct {
	FirstName string `form:"first_name" binding:"required"`
	LastName  string `form:"last_name"`
	Address   string `form:"address"`
	City      string `form:"city"`
	Telephone string `form:"telephone"`
	Pets      []int  `form:"pets[]"`
}

func Create(c *gin.Context) {
	var form OwnerForm

	if err := c.Bind(&form); err == nil {
		owner := &models.Owner{FirstName: form.FirstName, LastName: form.LastName}
		db.DB.Create(owner)
		logger.WithField("owner", owner).Infof("owner saved.")

		c.JSON(http.StatusOK, gin.H{"owner": owner})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}
