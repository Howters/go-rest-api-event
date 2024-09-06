package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/howters/golang/utils"
)

func Authenticate(context *gin.Context) {
	userToken := context.Request.Header.Get("Authorization")

	if userToken == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message" : "Not authorized",
		})
		return
	}

	userId, err := utils.VerifyToken(userToken)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message" : "Not authorized",
		})
		return
	}

	context.Set("userId", userId)

	context.Next()
}