package middleware

import (
	"net/http"

	"github.com/AlessandroCinque/GoAPI-Practice/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {

		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Not authorised"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Not authorised"})
		return
	}

	context.Set("userId", userId)

	context.Next()
}