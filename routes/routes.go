package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/miyabayt/gin-rest/routes/owners"
	"github.com/miyabayt/gin-rest/routes/pets"
	"github.com/miyabayt/gin-rest/routes/tasks"
)

func Use(router *gin.Engine) {
	// pets
	router.GET("/pets", pets.Index)
	router.GET("/pets/:id", pets.Show)
	router.POST("/pets", pets.Create)
	router.PUT("/pets/:id", pets.Update)

	// owners
	router.GET("/owners", owners.Index)
	router.POST("/owners", owners.Create)

	// tasks
	router.GET("/tasks", tasks.Index)
	router.GET("/tasks/:id", tasks.Show)
	router.POST("/tasks", tasks.Create)
	router.PUT("/tasks/:id", tasks.Update)
}
