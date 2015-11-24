package pets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miyabayt/gin-rest/db"
	"github.com/miyabayt/gin-rest/logger"
	"github.com/miyabayt/gin-rest/models"
)

type CreatePetForm struct {
	Name string `form:"name" binding:"required"`
	Age  int    `form:"age" binding:"required"`
}

func Create(c *gin.Context) {
	var form CreatePetForm

	if err := c.Bind(&form); err == nil {
		pet := &models.Pet{Name: form.Name, Age: form.Age}
		db.DB.Create(pet)
		logger.WithField("pet", pet).Infof("pet created.")

		c.JSON(http.StatusOK, gin.H{"pet": pet})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}
