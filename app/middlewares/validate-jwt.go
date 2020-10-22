package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-core-operations/app/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

// ValidateJwt : validates jwt
func ValidateJwt(context *gin.Context) {
	auth := context.GetHeader("Authorization")

	err := jwt.ValidateJWT(auth)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		context.Abort()
		return
	}
}
