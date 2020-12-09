package middlewares

import (
	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/gin-gonic/gin"
)

// FilterCanceledRequets : filter new requests
func FilterCanceledRequets(context *gin.Context) {
	transactions := context.Keys["transactions"].([]models.Transaction)
	var myTransactions []models.Transaction

	for _, transaction := range transactions {
		var shouldFilter bool
		for _, detail := range transaction.Details {
			if detail.Metadata == nil {
				continue
			}
			if detail.Status == "canceled" {
				shouldFilter = true
			}
		}

		if shouldFilter {
			continue
		}

		myTransactions = append(myTransactions, transaction)
	}

	context.Keys["transactions"] = myTransactions
}
