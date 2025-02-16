package api

import (
	movie_service "backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovie(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		return
	}

	movie := movie_service.GetById(id)

	context.JSON(http.StatusOK, movie)
}
