package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-core-operations/app/db/operationtypes"
	"github.com/Tsuryu/arreglapp-core-operations/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PostOperationtype : create a new operation type
func PostOperationtype(context *gin.Context) {
	operationType := models.OperationType{}
	context.ShouldBindBodyWith(&operationType, binding.JSON)

	err := operationtypes.Insert(&operationType)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create operation type", "operation_type": operationType})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, operationType)
}
