package middlewareutils

import (
	commonMiddlewares "github.com/Tsuryu/arreglapp-commons/app/middlewares"
	"github.com/Tsuryu/arreglapp-core-operations/app/middlewares"
	"github.com/gin-gonic/gin"
)

var middlewarelist = map[string][]gin.HandlerFunc{
	"health": {
		commonMiddlewares.Healthcheck,
	},
	"postOperationType": {
		middlewares.PostOperationtype,
	},
	"getOperationTypes": {
		middlewares.GetActive,
	},
	"postServiceRequest": {
		middlewares.PostServiceRequest,
	},
	// listServiceRequest
	"listServiceRequest": {
		middlewares.ListServiceRequest,
	},
}

// GetMiddlewares : get array of middlewares by name
func GetMiddlewares(name string) []gin.HandlerFunc {
	return middlewarelist[name]
}
