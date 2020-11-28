package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"github.com/gin-gonic/gin"
)

// ResponseOK : response ok from key
func ResponseOK(context *gin.Context) {
	response := utils.GetContextKey(context, "response")
	context.JSON(http.StatusOK, response)
}
