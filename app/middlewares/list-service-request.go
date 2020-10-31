package middlewares

import (
	"encoding/json"
	"net/http"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/Tsuryu/arreglapp-core-operations/app/models"
	"github.com/gin-gonic/gin"
)

// ListServiceRequest : get services by username
func ListServiceRequest(context *gin.Context) {
	operationTypes := context.Keys["operation_types"].([]models.OperationType)
	username := context.Query("username") // shortcut for c.Request.URL.Query().Get("username")
	transactions, err := service.GetTransactions(username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get services requests"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, marshallResponse(transactions, operationTypes))
}

func marshallResponse(transactions []commonModels.Transaction, operationTypes []models.OperationType) []models.ServiceRequest {
	serviceRequestList := []models.ServiceRequest{}
	for _, transaction := range transactions {
		if transaction.Reference != "service-request" {
			continue
		}

		metadata := transaction.Details[0].Metadata.(map[string]interface{})
		location := models.Location{}
		jsonString, _ := json.Marshal(metadata["location"])
		json.Unmarshal(jsonString, &location)

		operationType := models.OperationType{}
		for _, ot := range operationTypes {
			if ot.Name == metadata["type"].(string) {
				operationType = ot
			}
		}

		serviceRequestList = append(serviceRequestList, models.ServiceRequest{
			ID:            transaction.TraceID,
			Username:      transaction.Username,
			Type:          metadata["type"].(string),
			Title:         metadata["title"].(string),
			Description:   metadata["description"].(string),
			Location:      location,
			OperationType: operationType,
		})
	}

	return serviceRequestList
}
