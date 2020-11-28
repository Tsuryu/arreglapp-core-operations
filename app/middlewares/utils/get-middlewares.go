package middlewareutils

import (
	commonMiddlewares "github.com/Tsuryu/arreglapp-commons/app/middlewares"
	"github.com/Tsuryu/arreglapp-core-operations/app/middlewares"
	"github.com/gin-gonic/gin"
)

var middlewarelist = map[string][]gin.HandlerFunc{
	"health": {
		commonMiddlewares.Healthcheck,
	},
	"postOperationType": {
		middlewares.PostOperationtype,
	},
	"getOperationTypes": {
		middlewares.GetActive,
	},
	"postServiceRequest": {
		middlewares.PostServiceRequest,
	},
	"listServiceRequest": {
		middlewares.GetIcons,
		middlewares.ListServiceRequest,
		middlewares.MarshallServiceRequest,
		middlewares.ResponseOK,
	},
	"searchServiceRequest": { // search new request
		middlewares.GetIcons,
		middlewares.SearchServiceRequest,
		middlewares.FilterMyServiceRequest,
		middlewares.FilterOngoingServiceRequest,
		middlewares.MarshallServiceRequest,
		middlewares.ResponseOK,
	},
	"initChatServiceRequest": {
		middlewares.InitChat,
		middlewares.GetTransaction,
		middlewares.GetIcons,
		middlewares.MarshallServiceRequest,
		middlewares.NotifyNewChat,
	},
	"professionalListServiceRequest": { // ongoing
		middlewares.GetIcons,
		middlewares.SearchServiceRequest,
		middlewares.FilterMyServiceRequest,
		middlewares.FilterNewRequests,
		middlewares.MarshallProfessionalServiceRequest,
		middlewares.ResponseOK,
	},
	"confirmProfessional": {
		middlewares.ConfirmProfessional,
		middlewares.GetIcons,
		middlewares.GetTransaction,
		middlewares.MarshallProfessionalServiceRequest,
		middlewares.NotifyConfirmProfessional,
	},
	"postBudget": {
		middlewares.GetTransaction,
		middlewares.PostBudget,
		middlewares.GetTransaction,
		middlewares.GetIcons,
		middlewares.MarshallServiceRequest,
		middlewares.NotifyNewBudget,
	},
}

// GetMiddlewares : get array of middlewares by name
func GetMiddlewares(name string) []gin.HandlerFunc {
	return middlewarelist[name]
}
