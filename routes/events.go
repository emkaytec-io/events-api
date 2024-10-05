package routes

import (
	"net/http"
	"strconv"

	"emkaytec.io/events-api/v2/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events, please try again later.",
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID.",
		})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event, please try again later.",
		})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	event.UserID = context.GetInt64("userId")

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event, please try again later.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created.",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID.",
		})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event, please try again later.",
		})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update event.",
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event, please try again later.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated.",
		"event":   updatedEvent,
	})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID.",
		})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event, please try again later.",
		})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to delete event.",
		})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event, please try again later.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted.",
	})
}
