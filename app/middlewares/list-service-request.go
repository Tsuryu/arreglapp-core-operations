package middlewares

import (
	"net/http"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/gin-gonic/gin"
)

// ListServiceRequest : get services by username
func ListServiceRequest(context *gin.Context) {
	username := context.Query("username") // shortcut for c.Request.URL.Query().Get("username")
	transactions, err := service.GetTransactions(commonModels.Transaction{Username: username})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get services requests"})
		context.Abort()
		return
	}

	context.Keys["transactions"] = transactions
}
