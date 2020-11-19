package middlewares

import (
	"fmt"
	"net/http"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/Tsuryu/arreglapp-core-operations/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PostServiceRequest : init service request ransaction
func PostServiceRequest(context *gin.Context) {
	claim := context.Keys["claims"].(*commonModels.Claim)
	fmt.Println(claim)
	serviceRequest := models.ServiceRequest{}
	context.ShouldBindBodyWith(&serviceRequest, binding.JSON)
	serviceRequest.UserContactInfo = models.UserContactInfo{
		Firstname: claim.FirstName,
		Lastname:  claim.LastName,
		Phone:     claim.Phone,
	}

	transaction := commonModels.Transaction{}
	transactionDetail := commonModels.TransactionDetail{}
	transactionDetail.Status = "created"

	transactionDetail.Metadata = serviceRequest
	transaction.Reference = "service-request"
	transaction.Username = serviceRequest.Username
	transaction.Details = append(transaction.Details, transactionDetail)

	err := service.NewTransaction(&transaction, true)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to init transaction"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"trace_id": transaction.TraceID})
}
