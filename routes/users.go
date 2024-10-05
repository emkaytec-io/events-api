package routes

import (
	"net/http"

	"emkaytec.io/events-api/v2/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user, please try again later.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created.",
		"user":    user,
	})

}
