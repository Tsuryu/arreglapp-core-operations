package middlewares

import (
	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/gin-gonic/gin"
)

// FilterExceptHistory : filter new requests
func FilterExceptHistory(context *gin.Context) {
	transactions := context.Keys["transactions"].([]models.Transaction)
	var myTransactions []models.Transaction
	var shouldFilter bool

	for _, transaction := range transactions {
		shouldFilter = true
		for _, detail := range transaction.Details {
			if detail.Metadata == nil {
				continue
			}
			if detail.Status == "canceled" || detail.Status == "finished" {
				shouldFilter = false
			}
		}

		if shouldFilter {
			continue
		}

		myTransactions = append(myTransactions, transaction)
	}

	context.Keys["transactions"] = myTransactions
}
