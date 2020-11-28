package middlewares

import (
	"fmt"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"github.com/Tsuryu/arreglapp-core-operations/app/models"
	"github.com/gin-gonic/gin"
)

// NotifyNewBudget : send a push notification to notify that the provider filled the budget
func NotifyNewBudget(context *gin.Context) {
	requests := utils.GetContextKey(context, "response").([]models.ServiceRequest)
	username := utils.GetContextKey(context, "username").(string)
	id, err := service.GetPushNotificationID(username)
	if err != nil {
		fmt.Println("Failed to get push notification id")
		return
	}

	err = service.SendPushNotification(commonModels.PushNotification{
		To: id,
		Notification: commonModels.NotificationInfo{
			Title: "Actualizacion de presupuesto",
			Body:  "El proveedor te ha enviado el presupuesto para tu pedido",
		},
		Data: struct {
			JobRequest models.ServiceRequest `json:"job_request"`
		}{
			JobRequest: requests[0],
		},
	})
	if err != nil {
		fmt.Println("Failed to send init chat push notification")
	}
}
