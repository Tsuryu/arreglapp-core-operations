package routers

import (
	"os"

	"github.com/Tsuryu/arreglapp-core-operations/app/middlewares"
	middlewareutils "github.com/Tsuryu/arreglapp-core-operations/app/middlewares/utils"
	"github.com/gin-gonic/gin"
)

// Router : app routes
func Router() {
	router := gin.Default()
	router.Use(middlewares.ValidateJwt)
	router.GET("/health", middlewareutils.GetMiddlewares("health")...)
	router.POST("/operation-type", middlewareutils.GetMiddlewares("postOperationType")...)
	router.GET("/operation-type", middlewareutils.GetMiddlewares("getOperationTypes")...)

	router.Run(":" + os.Getenv("APP_PORT"))
}
