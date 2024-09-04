package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/howters/golang/models"
	"github.com/howters/golang/utils"
)

func signUp(context *gin.Context) {
	var user models.User
	//Automatically populate the request body to the event variable
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message" : "Could not parse request data.",
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message" : "Could not save user.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message" : "User created successfully.",
	})
}

func logIn(context *gin.Context) {
	var user models.User
	//Automatically populate the request body to the event variable
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message" : "Could not parse request data.",
		})
		return
	}

	err = user.VerifyCredentials()


	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message" : "Could not authenticate user.",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message" : "Could not authenticate user.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message" : "User created successfully.",
		"token" : token,
	})
}