package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/gin-gonic/gin"
)

// PayTransactionFee : initialize chat with the client
func PayTransactionFee(context *gin.Context) {
	transactionDetail := models.TransactionDetail{}
	body := struct {
		Image string `json:"image"`
	}{}

	context.ShouldBind(&body)

	traceID := context.GetHeader("trace-id")
	transactionDetail.Status = "transaction-fee-payment"
	transactionDetail.Metadata = struct {
		Image string
	}{
		Image: body.Image,
	}

	result, err := service.AddTransactionDetail(transactionDetail, traceID, "")
	if err != nil || !result {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to pay transaction fee"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})
}
