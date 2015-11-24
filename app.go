package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/miyabayt/gin-rest/config"
	"github.com/miyabayt/gin-rest/logger"
	"github.com/miyabayt/gin-rest/routes"
)

var (
	listenPort = ":3001"
)

func startHttpServer() {
	router := gin.New() // don't need to use Logger, if we use apache/nginx.
	router.Use(gin.Recovery())
	routes.Use(router)

	if gin.Mode() == gin.ReleaseMode {
		logger.Infof("GIN Listening and serving HTTP on %s", listenPort)
		endless.ListenAndServe(listenPort, router) // use endless for graceful restart/stop.
	} else {
		router.Run(listenPort)
	}
}

func main() {
	appName, _ := config.String("app.name")
	logger.Infof("starting %s", appName)

	startHttpServer()
}

func init() {
	numCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu) // until "go 1.5"

	logger.Infof("GOMAXPROCS set to %d.", numCpu)

	if envPort := os.Getenv("PORT"); 0 < len(envPort) {
		listenPort = fmt.Sprintf(":%s", envPort)
		logger.Infof("port set to %s.", listenPort)
	}
}
