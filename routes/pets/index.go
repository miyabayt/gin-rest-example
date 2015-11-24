package pets

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miyabayt/gin-rest/config"
	"github.com/miyabayt/gin-rest/db"
	"github.com/miyabayt/gin-rest/logger"
	"github.com/miyabayt/gin-rest/models"
)

func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		page = 1
	}

	perpage, _ := config.Int("app.perpage")
	offset := (page - 1) * perpage
	if page < 2 {
		offset = 0
	}

	var count int
	var pets []models.Pet
	db.DB.Find(&pets).Count(&count)
	db.DB.Limit(perpage).Offset(offset).Find(&pets)
	logger.Infof("%d pets found.", len(pets))

	data := map[string]interface{}{"pets": pets}
	c.JSON(http.StatusOK, gin.H{"data": data, "page": page, "perpage": perpage, "total": count})
}
