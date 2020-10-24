package routers

import (
	"os"

	commonMiddlewares "github.com/Tsuryu/arreglapp-commons/app/middlewares"
	middlewareutils "github.com/Tsuryu/arreglapp-core-operations/app/middlewares/utils"
	"github.com/gin-gonic/gin"
)

// Router : app routes
func Router() {
	router := gin.Default()
	router.Use(commonMiddlewares.ValidateJwt)
	router.GET("/health", middlewareutils.GetMiddlewares("health")...)
	router.POST("/operation-type", middlewareutils.GetMiddlewares("postOperationType")...)
	router.GET("/operation-type", middlewareutils.GetMiddlewares("getOperationTypes")...)
	router.POST("/service-request", middlewareutils.GetMiddlewares("postServiceRequest")...)
	router.GET("/service-request/list", middlewareutils.GetMiddlewares("listServiceRequest")...)

	router.Run(":" + os.Getenv("APP_PORT"))
}
