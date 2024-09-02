package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/howters/golang/models"
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