package middlewares

import (
	"encoding/json"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"github.com/gin-gonic/gin"
)

// GetServiceRequestProviderFromTransaction : get provider username to notify about new transaction fee
func GetServiceRequestProviderFromTransaction(context *gin.Context) {
	transactions := utils.GetContextKey(context, "transactions").([]commonModels.Transaction)
	for _, transaction := range transactions {
		for _, detail := range transaction.Details {
			if detail.Status == "budget" {
				metadata := struct {
					Username string `json:"Username"`
				}{}
				chatJSONString, _ := json.Marshal(detail.Metadata)
				json.Unmarshal(chatJSONString, &metadata)

				utils.AddContextKey(context, "username", metadata.Username)
			}
		}
	}
}
