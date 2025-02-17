package api

import (
	"backend/repository"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie_controller struct {
	repository repository.MovieRepository
}

func NewMovieController(repository repository.MovieRepository) *Movie_controller {
	return &Movie_controller{
		repository: repository,
	}
}

func (mc *Movie_controller) GetMovieById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		return
	}
	
	movie := movie_service.NewMovieService(mc.repository).GetById(id)

	c.JSON(http.StatusOK, movie)
}
