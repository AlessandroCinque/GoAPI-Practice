package routes

import(
	"fmt"
	"net/http"
	"strconv"
	"github.com/AlessandroCinque/GoAPI-Practice/models"
	"github.com/gin-gonic/gin"
)

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"),10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event Id"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Culd not fetch events. 500"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Culd not fetch events. 500"})
	}

	//The gin package will automatically turn the structs into JSONs
	context.JSON(http.StatusOK, events)
}


func CreateEvent(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request", "event": event})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Culd not fetch events. 500"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}