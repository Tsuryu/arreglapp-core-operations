package middlewares

import (
	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/gin-gonic/gin"
)

// FilterNewRequests : filter new requests
func FilterNewRequests(context *gin.Context) {
	transactions := context.Keys["transactions"].([]models.Transaction)
	claim := context.Keys["claims"].(*models.Claim)
	var myTransactions []models.Transaction
	var shouldFilter bool

	for _, transaction := range transactions {
		shouldFilter = true
		for _, detail := range transaction.Details {
			if detail.Metadata == nil {
				continue
			}
			detailMetadata := detail.Metadata.(map[string]interface{})
			if detail.Status == "new-chat" && detailMetadata["Username"] == claim.Username {
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
