package owners

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
	var owners []models.Owner
	db.DB.Find(&owners).Count(&count)
	db.DB.Limit(perpage).Offset(offset).Find(&owners)
	logger.Infof("%d owners found.", len(owners))

	data := map[string]interface{}{"owners": owners}
	c.JSON(http.StatusOK, gin.H{"data": data, "page": page, "perpage": perpage, "total": count})
}
