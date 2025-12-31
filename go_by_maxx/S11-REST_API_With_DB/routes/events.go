package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest_api_with_db/models"
	"github.com/gin-gonic/gin"
)

// whenever this fn get attached to a route gin passes context to it
func GetEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events. Try again later",
		})

		return
	}

	// gin.H is a shortcut for map[string]interface{}
	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully fetched events",
		"events":  events,
	})
}

func CreateEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event) // gin will scan the request body and put it in the event variable

	fmt.Printf("Body : %v", event)
	fmt.Printf("Err : %v", err)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event. Try again later",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event":   event,
	})
}

func GetEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event. Try again later",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully fetched event",
		"event":   event,
	})
}

func UpdateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	userId := context.GetInt64("userId")
	// Check if event exists before updating
	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update this event",
		})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	updatedEvent.Id = id

	err = updatedEvent.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event. Try again later",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated event",
		"event":   updatedEvent,
	})
}

func DeleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	userId := context.GetInt64("userId")
	// Check if event exists before deleting
	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to delete this event",
		})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event. Try again later",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted event",
	})
}
