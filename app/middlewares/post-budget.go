package middlewares

import (
	"net/http"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"github.com/Tsuryu/arreglapp-core-operations/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PostBudget : creates a budget by the professional
func PostBudget(context *gin.Context) {
	claim := context.Keys["claims"].(*commonModels.Claim)
	transactionDetail := commonModels.TransactionDetail{}
	budget := models.Budget{}
	context.ShouldBindBodyWith(&budget, binding.JSON)
	budget.Username = claim.Username

	traceID := context.GetHeader("trace-id")
	transactionDetail.Status = "budget"
	transactionDetail.Metadata = budget

	result, err := service.AddTransactionDetail(transactionDetail, traceID, "")
	if err != nil || !result {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create budget"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})

	transactions := context.Keys["transactions"].([]commonModels.Transaction)
	utils.AddContextKey(context, "username", transactions[0].Username)
}
