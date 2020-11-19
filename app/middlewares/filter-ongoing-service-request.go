package middlewares

import (
	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/gin-gonic/gin"
)

// FilterOngoingServiceRequest : filter ongoing requests
func FilterOngoingServiceRequest(context *gin.Context) {
	transactions := context.Keys["transactions"].([]models.Transaction)
	claim := context.Keys["claims"].(*models.Claim)
	var ongoingTransactions []models.Transaction
	var shouldContinue bool

	for _, transaction := range transactions {
		if transaction.Username != claim.Username {

			for _, detail := range transaction.Details {
				detailMetadata := detail.Metadata.(map[string]interface{})
				if detail.Status == "new-chat" && detailMetadata["Username"] == claim.Username {
					shouldContinue = true
				}
			}

			if shouldContinue {
				shouldContinue = false
				continue
			}

			ongoingTransactions = append(ongoingTransactions, transaction)
		}
	}

	context.Keys["transactions"] = ongoingTransactions
}
