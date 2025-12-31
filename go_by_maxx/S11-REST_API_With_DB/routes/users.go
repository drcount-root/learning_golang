package routes

import (
	"fmt"
	"net/http"

	"example.com/rest_api_with_db/models"
	"example.com/rest_api_with_db/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	fmt.Printf("Body : %v", &user)
	fmt.Printf("Err : %v", err)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
			"error":   err,
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not sign up. Try again later",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Successfully signed up",
		"user":    user,
	})
}

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	fmt.Printf("Body : %v", &user)
	fmt.Printf("Err : %v", err)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user.",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token. Try again later",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"token":   token,
	})
}
