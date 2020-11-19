package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/gin-gonic/gin"
)

// InitChat : initialize chat with the client
func InitChat(context *gin.Context) {
	transactionDetail := models.TransactionDetail{}
	// context.ShouldBindBodyWith(&transactionDetail, binding.JSON)
	claim := context.Keys["claims"].(*models.Claim)

	traceID := context.GetHeader("trace-id")
	transactionDetail.Status = "new-chat"
	transactionDetail.Metadata = struct {
		FirstName string
		LastName  string
		Phone     string
		Username  string
	}{
		claim.FirstName,
		claim.LastName,
		claim.Phone,
		claim.Username,
	}

	result, err := service.AddTransactionDetail(transactionDetail, traceID, "")
	if err != nil || !result {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to init transaction"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})
}
