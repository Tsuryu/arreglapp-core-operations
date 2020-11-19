package middlewares

import (
	"encoding/json"
	"net/http"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-core-operations/app/models"
	"github.com/gin-gonic/gin"
)

// MarshallServiceRequest : mix services request with their icons
func MarshallServiceRequest(context *gin.Context) {
	operationTypes := context.Keys["operation_types"].([]models.OperationType)
	transactions := context.Keys["transactions"].([]commonModels.Transaction)
	// claim := context.Keys["claims"].(*commonModels.Claim)
	// var shouldContinue bool

	serviceRequestList := []models.ServiceRequest{}
	for _, transaction := range transactions {
		// if transaction.Reference != "service-request" || claim.Username != transaction.Username {
		if transaction.Reference != "service-request" {
			continue
		}

		// for _, detail := range transaction.Details {
		// 	detailMetadata := detail.Metadata.(map[string]interface{})
		// 	if detailMetadata["Phone"] == claim.Phone {
		// 		shouldContinue = true
		// 	}
		// }

		// if shouldContinue {
		// 	shouldContinue = false
		// 	continue
		// }

		metadata := transaction.Details[0].Metadata.(map[string]interface{})
		location := models.Location{}
		locationJSONString, _ := json.Marshal(metadata["location"])
		json.Unmarshal(locationJSONString, &location)

		operationType := models.OperationType{}
		for _, ot := range operationTypes {
			if ot.Name == metadata["type"].(string) {
				operationType = ot
			}
		}

		confirmedUser := models.UserContactInfo{}
		for _, detail := range transaction.Details {
			if detail.Status == "chosen-professional" {
				chatJSONString, _ := json.Marshal(detail.Metadata)
				json.Unmarshal(chatJSONString, &confirmedUser)
			}
		}

		chats := make([]models.UserContactInfo, 0)
		for _, detail := range transaction.Details {
			if detail.Status == "new-chat" {
				chat := models.UserContactInfo{}
				chatJSONString, _ := json.Marshal(detail.Metadata)
				json.Unmarshal(chatJSONString, &chat)

				if confirmedUser.Username == chat.Username {
					chat.Confirmed = true
				}

				chats = append(chats, chat)
			}
		}

		userContactInfo := models.UserContactInfo{}
		userContactInfoJSONString, _ := json.Marshal(metadata["user_contact_info"])
		json.Unmarshal(userContactInfoJSONString, &userContactInfo)

		serviceRequestList = append(serviceRequestList, models.ServiceRequest{
			ID:              transaction.TraceID,
			Username:        transaction.Username,
			Type:            metadata["type"].(string),
			Title:           metadata["title"].(string),
			Description:     metadata["description"].(string),
			UserContactInfo: userContactInfo,
			Location:        location,
			OperationType:   operationType,
			Chats:           chats,
		})
	}

	context.JSON(http.StatusOK, serviceRequestList)
}
