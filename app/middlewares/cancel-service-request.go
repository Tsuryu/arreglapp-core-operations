package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/gin-gonic/gin"
)

// CancelServiceRequest : cancel a service request by user
func CancelServiceRequest(context *gin.Context) {
	transactionDetail := models.TransactionDetail{}
	claim := context.Keys["claims"].(*models.Claim)
	uriParams := struct {
		ID string `uri:"id"`
	}{}

	context.ShouldBindUri(&uriParams)

	traceID := uriParams.ID
	transactionDetail.Status = "canceled"
	transactionDetail.Metadata = struct {
		Username string
	}{
		claim.Username,
	}

	result, err := service.AddTransactionDetail(transactionDetail, traceID, "")
	if err != nil || !result {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to cancel service request"})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}
