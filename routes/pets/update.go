package pets

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miyabayt/gin-rest/db"
	"github.com/miyabayt/gin-rest/logger"
	"github.com/miyabayt/gin-rest/models"
)

type UpdatePetForm struct {
	Name    string `form:"name" binding:"required"`
	Age     int    `form:"age" binding:"required"`
	OwnerId int64  `form:"owner_id"`
}

func Update(c *gin.Context) {
	var form UpdatePetForm

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

	}

	if err := c.Bind(&form); err == nil {
		pet := &models.Pet{OwnerId: form.OwnerId}
		db.DB.First(pet, id).Update(pet)
		logger.WithField("pet", pet).Infof("pet saved.")

		c.JSON(http.StatusOK, gin.H{"pet": pet})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}
