package routes

import (
	"fmt"
	"net/http"

	"github.com/AlessandroCinque/GoAPI-Practice/modelsWithDBQueries"
	"github.com/AlessandroCinque/GoAPI-Practice/utils"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {

	var user modelsWithDBQueries.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request", "user": user})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not save user", "user": user})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"User created successfullly"})
}


func Login(context *gin.Context) {

	var user modelsWithDBQueries.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request", "user": user})
		return
	}

	err = user.ValidateCredentials()
	
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	
	context.JSON(http.StatusOK,  gin.H{"message": "Login successful", "token": token})

}
