package middlewares

import (
	"encoding/json"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"github.com/Tsuryu/arreglapp-core-operations/app/models"
	"github.com/gin-gonic/gin"
)

// MarshallProfessionalServiceRequest : mix services request with their icons
func MarshallProfessionalServiceRequest(context *gin.Context) {
	operationTypes := context.Keys["operation_types"].([]models.OperationType)
	transactions := context.Keys["transactions"].([]commonModels.Transaction)
	claim := context.Keys["claims"].(*commonModels.Claim)
	var confirmed, canceled bool

	serviceRequestList := []models.ServiceRequest{}
	for _, transaction := range transactions {
		if transaction.Reference != "service-request" {
			continue
		}

		var payed, transactionFeePayed bool
		budget := models.Budget{}
		confirmed = false

		for _, detail := range transaction.Details {
			if detail.Status == "service-payment" {
				payed = true
			}

			if detail.Status == "transaction-fee-payment" {
				transactionFeePayed = true
			}

			if detail.Metadata == nil {
				continue
			}
			detailMetadata := detail.Metadata.(map[string]interface{})
			if detail.Status == "chosen-professional" && detailMetadata["Username"] == claim.Username {
				confirmed = true
			}

			if detail.Status == "budget" {
				jsonString, _ := json.Marshal(detailMetadata)
				json.Unmarshal(jsonString, &budget)
			}
			if detail.Status == "canceled" {
				canceled = true
			}
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

		userContactInfo := models.UserContactInfo{}
		userContactInfoJSONString, _ := json.Marshal(metadata["user_contact_info"])
		json.Unmarshal(userContactInfoJSONString, &userContactInfo)
		userContactInfo.Confirmed = confirmed

		serviceRequest := models.ServiceRequest{
			ID:                  transaction.TraceID,
			Username:            transaction.Username,
			Type:                metadata["type"].(string),
			Title:               metadata["title"].(string),
			Description:         metadata["description"].(string),
			UserContactInfo:     userContactInfo,
			Location:            location,
			OperationType:       operationType,
			Payed:               payed,
			TransactionFeePayed: transactionFeePayed,
			Canceled:            canceled,
		}

		if budget.Username != "" {
			serviceRequest.Budget = &budget
		}

		serviceRequestList = append(serviceRequestList, serviceRequest)
	}

	utils.AddContextKey(context, "response", serviceRequestList)
}
