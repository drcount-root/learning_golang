package main

import (
	"fmt"
	"net/http"

	"example.com/rest_api_with_db/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

// whenever this fn get attached to a route gin passes context to it
func getEvents(context *gin.Context) {

	events := models.GetAllEvents()

	// gin.H is a shortcut for map[string]interface{}
	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully fetched events",
		"events":  events,
	})
}

func createEvent(context *gin.Context) {
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

	event.Id = 1
	event.UserId = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event":   event,
	})
}
