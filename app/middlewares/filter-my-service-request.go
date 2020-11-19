package middlewares

import (
	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/gin-gonic/gin"
)

// FilterMyServiceRequest : filter my own service requests
func FilterMyServiceRequest(context *gin.Context) {
	transactions := context.Keys["transactions"].([]models.Transaction)
	claim := context.Keys["claims"].(*models.Claim)
	var myTransactions []models.Transaction

	for _, transaction := range transactions {
		if transaction.Username != claim.Username {
			myTransactions = append(myTransactions, transaction)
		}
	}

	context.Keys["transactions"] = myTransactions
}
