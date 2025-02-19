package api

import (
	"backend/errors"
	"backend/model"
	"backend/repository"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie_controller struct {
	repository repository.MovieRepository
}

func (mc *Movie_controller) InsertNewMovie(c *gin.Context) {
	var movie_data model.Movie

	if err := c.ShouldBindBodyWithJSON(&movie_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "required data are either missing or malformed."})
		return
	}

	id, err := services.NewMovieService(mc.repository).InsertNewMovie(movie_data)

	if err != nil {
		if internalErr, ok := err.(*errors.InternalError); ok {
			c.JSON(internalErr.Code, gin.H{"error": internalErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "resource created",
		"resource_id": id,
	})
}

func (mc *Movie_controller) GetMovieById(c *gin.Context) {
	movie, err := services.NewMovieService(mc.repository).GetMovieById(c.Param("id"))

	if err != nil {
		if internalErr, ok := err.(*errors.InternalError); ok {
			c.JSON(internalErr.Code, gin.H{"error": internalErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (mc *Movie_controller) GetListAllMovies(c *gin.Context) {
	movie, err := services.NewMovieService(mc.repository).GetListAllMovies()

	if err != nil {
		if internalErr, ok := err.(*errors.InternalError); ok {
			c.JSON(internalErr.Code, gin.H{"error": internalErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (mc *Movie_controller) DeleteMovieById(c *gin.Context) {
	err := services.NewMovieService(mc.repository).DeleteMovieById(c.Param("id"))

	if err != nil {
		if internalErr, ok := err.(*errors.InternalError); ok {
			c.JSON(internalErr.Code, gin.H{"error": internalErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "resource deleted succesfully",
	})
}

func (mc *Movie_controller)GetMoviesByPublisherId(c *gin.Context) {
	movies, err := services.NewMovieService(mc.repository).GetMoviesByPublisherId(c.Param("id"))

	if err != nil {
		if internalErr, ok := err.(*errors.InternalError); ok {
			c.JSON(internalErr.Code, gin.H{"error": internalErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	
	c.JSON(http.StatusOK, movies)
}

func NewMovieController(repository repository.MovieRepository) *Movie_controller {
	return &Movie_controller{
		repository: repository,
	}
}