package middlewares

import (
	"net/http"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/gin-gonic/gin"
)

// SearchServiceRequest : fetch request beeing a provider
func SearchServiceRequest(context *gin.Context) {
	transactions, err := service.GetTransactions(commonModels.Transaction{Reference: "service-request"})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get services requests"})
		context.Abort()
		return
	}

	context.Keys["transactions"] = transactions
}
