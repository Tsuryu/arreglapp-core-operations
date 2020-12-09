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
		middlewares.FilterCanceledRequets,
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
		middlewares.FilterCanceledRequets,
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
	"payServiceRequest": {
		middlewares.GetTransaction,
		middlewares.PayServiceRequest,
		middlewares.GetServiceRequestProviderFromTransaction,
		middlewares.GetTransaction,
		middlewares.GetIcons,
		middlewares.MarshallProfessionalServiceRequest,
		middlewares.NotifyNewTransactionFee,
	},
	"payTransactionFee": {
		middlewares.PayTransactionFee,
	},
	"supportRequest": {
		middlewares.PostSupportRequest,
	},
	"cancelServiceRequest": {
		middlewares.CancelServiceRequest,
	},
	"professionalHistoryServiceRequest": {
		middlewares.GetIcons,
		middlewares.SearchServiceRequest,
		middlewares.FilterMyServiceRequest,
		middlewares.FilterExceptHistory,
		middlewares.MarshallProfessionalServiceRequest,
		middlewares.ResponseOK,
	},
}

// GetMiddlewares : get array of middlewares by name
func GetMiddlewares(name string) []gin.HandlerFunc {
	return middlewarelist[name]
}
