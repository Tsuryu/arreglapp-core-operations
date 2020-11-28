package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"github.com/gin-gonic/gin"
)

// GetTransaction : get transaction from api
func GetTransaction(context *gin.Context) {
	traceID := context.GetHeader("trace-id")
	transaction, err := service.GetTransaction(traceID)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	transactions := []models.Transaction{}
	transactions = append(transactions, transaction)

	utils.AddContextKey(context, "transactions", transactions)
}
