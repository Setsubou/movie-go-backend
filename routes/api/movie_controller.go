package api

import (
	"backend/repository"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewMovieController(repository repository.MovieRepository) *Movie_controller {
	return &Movie_controller{
		repository: repository,
	}
}

type Movie_controller struct {
	repository repository.MovieRepository
}

func (mc *Movie_controller) GetMovieById(c *gin.Context) {
	movie := services.NewMovieService(mc.repository).GetMovieById(c.Param("id"))

	c.JSON(http.StatusOK, movie)
}

func (mc *Movie_controller) GetListAllMovies(c *gin.Context) {
	movie := services.NewMovieService(mc.repository).GetListAllMovies()

	c.JSON(http.StatusOK, movie)
}
