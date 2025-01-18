package routes

import (
	"event-management/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch error",
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id.",
		})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch event",
		})
		return
	}
	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {

	// var userId int64
	// var err error

	userId := context.GetInt64("userId")

	var event models.Event
	err := context.ShouldBind(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data.",
		})
		return
	}

	// event.Id = 1
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created !",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id.",
		})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch event",
		})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Your are not authorized to make changes",
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBind(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data.",
		})
		return
	}
	updatedEvent.Id = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not update event",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}

func deletEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userId := context.GetInt64("userId")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id.",
		})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch event",
		})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Your are not authorized to make changes",
		})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not delete event",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
