package middlewares

import (
	"net/http"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"github.com/gin-gonic/gin"
)

// ConfirmProfessional : choose an professional to carry out the fix
func ConfirmProfessional(context *gin.Context) {
	transactionDetail := commonModels.TransactionDetail{}
	username := context.Param("id")
	traceID := context.GetHeader("trace-id")

	utils.AddContextKey(context, "username", username)
	transactionDetail.Status = "chosen-professional"
	transactionDetail.Metadata = struct {
		Username string
	}{
		username,
	}

	result, err := service.AddTransactionDetail(transactionDetail, traceID, "")
	if err != nil || !result {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to chose professional"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})

	utils.AddContextKey(context, "claims", &commonModels.Claim{
		Username: username,
	})
}
