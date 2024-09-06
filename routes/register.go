package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/howters/golang/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err:=strconv.ParseInt(context.Param("id"),10 , 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message" : "Could not parse event id.",
		})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message" : "Could not find event",
		})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message" : "Could not register event",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message" : "Event registered",
	})
	return

}


func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err:=strconv.ParseInt(context.Param("id"),10 , 64)

	event, err := models.GetEventByID(eventId)

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H {
			"message" : "Could not cancel registration",
		})

		return;
	}

	context.JSON(http.StatusCreated, gin.H {
		"message" : "Created",
	})


}