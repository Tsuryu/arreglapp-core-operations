package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-core-operations/app/db/operationtypes"
	"github.com/gin-gonic/gin"
)

// GetActive : create a new operation type
func GetActive(context *gin.Context) {
	result, err := operationtypes.Find()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get operation types"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, result)
}
