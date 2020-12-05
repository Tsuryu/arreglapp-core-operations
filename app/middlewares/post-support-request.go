package middlewares

import (
	"net/http"
	"time"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	supportrequets "github.com/Tsuryu/arreglapp-core-operations/app/db/supportrequests"
	"github.com/Tsuryu/arreglapp-core-operations/app/models"
	"github.com/gin-gonic/gin"
)

// PostSupportRequest : create a new support request
func PostSupportRequest(context *gin.Context) {
	claim := context.Keys["claims"].(*commonModels.Claim)

	body := struct {
		Description string `json:"description"`
	}{}

	context.ShouldBind(&body)

	supportRequest := models.SupportRequest{
		Username:    claim.Username,
		Description: body.Description,
		Date:        time.Now(),
	}

	err := supportrequets.Insert(supportRequest)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create request ticket"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})
}
