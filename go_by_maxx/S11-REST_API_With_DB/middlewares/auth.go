package middlewares

import (
	"net/http"
	// "strings"

	"example.com/rest_api_with_db/models"
	"example.com/rest_api_with_db/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	// // Handle both "Bearer <token>" and raw token formats
	// if strings.HasPrefix(token, "Bearer ") {
	// 	token = strings.TrimPrefix(token, "Bearer ")
	// }

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"error":   err.Error(),
		})
		return
	}

	// Verify that the user from the token actually exists in the database
	userExists, err := models.UserExists(userId)

	if err != nil || !userExists {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized - User does not exist",
		})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
