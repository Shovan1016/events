package routes

import (
	"event-management/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id.",
		})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "event not exist",
		})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not register for the event",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "you are register for the event successfully",
	})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, _ := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event

	event.Id = eventId

	err := event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not cancel registration",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "cancelled",
	})
}
