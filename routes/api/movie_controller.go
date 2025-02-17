package api

import (
	"backend/repository"
	"backend/services"
	"net/http"

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
	movie := movie_service.NewMovieService(mc.repository).GetById(c.Param("id"))

	c.JSON(http.StatusOK, movie)
}
