package middlewares

import (
	"log"

	"github.com/Tsuryu/arreglapp-core-operations/app/db/operationtypes"
	"github.com/gin-gonic/gin"
)

// GetIcons : get operation types icons to match afterwards
func GetIcons(context *gin.Context) {
	result, err := operationtypes.Find()
	if err != nil {
		log.Println(err)
		return
	}

	context.Keys["operation_types"] = result
}
