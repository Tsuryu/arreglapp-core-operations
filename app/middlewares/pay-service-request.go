package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"github.com/gin-gonic/gin"
)

// PayServiceRequest : client confirms that he payed the service to the provider
func PayServiceRequest(context *gin.Context) {
	transactions := utils.GetContextKey(context, "transactions").([]models.Transaction)
	claim := context.Keys["claims"].(*models.Claim)

	if transactions[0].Username != claim.Username {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid user transaction"})
		return
	}

	traceID := context.GetHeader("trace-id")
	transactionDetail := models.TransactionDetail{}
	transactionDetail.Status = "service-payment"

	result, err := service.AddTransactionDetail(transactionDetail, traceID, "")
	if err != nil || !result {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to init transaction"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})
}
