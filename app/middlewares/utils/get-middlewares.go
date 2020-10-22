package middlewareutils

import (
	"github.com/Tsuryu/arreglapp-core-operations/app/middlewares"
	"github.com/gin-gonic/gin"
)

var middlewarelist = map[string][]gin.HandlerFunc{
	"health": {
		middlewares.Healthcheck,
	},
	"postOperationType": {
		middlewares.PostOperationtype,
	},
	"getOperationTypes": {
		middlewares.GetActive,
	},
}

// GetMiddlewares : get array of middlewares by name
func GetMiddlewares(name string) []gin.HandlerFunc {
	return middlewarelist[name]
}
