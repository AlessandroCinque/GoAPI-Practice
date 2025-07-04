package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlessandroCinque/GoAPI-Practice/modelsWithDBQueries"
	"github.com/gin-gonic/gin"
)

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"),10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event Id"})
		return
	}

	event, err := modelsWithDBQueries.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Culd not fetch events. 500"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func GetEvents(context *gin.Context) {
	events, err := modelsWithDBQueries.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Culd not fetch events. 500"})
	}

	//The gin package will automatically turn the structs into JSONs
	context.JSON(http.StatusOK, events)
}


func CreateEvent(context *gin.Context) {

	//middleware.Authenticate(context)

	var event modelsWithDBQueries.Event

	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request", "event": event})
		return
	}

	
	event.UserID = context.GetInt64("userId")

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Culd not fetch events. 500"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"),10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event Id"})
		return
	}

	userId := context.GetInt64("userId")
	event , err := modelsWithDBQueries.GetEventByID(eventId)

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You don't have the permission to chnage this event"})

		return
	}
	

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch the event"})
		return
	}



	var updatedEvent modelsWithDBQueries.Event
	err = context.ShouldBindBodyWithJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event Id"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not update the event"})
		return
	}

	
	context.JSON(http.StatusOK, gin.H{"message": "Event update successfully"})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"),10, 64)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event Id"})
		return
	}

	event, err := modelsWithDBQueries.GetEventByID(eventId)

	userId := context.GetInt64("userId")

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You don't have the permission to chnage this event"})

		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch the event"})
		return
	}

	event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}