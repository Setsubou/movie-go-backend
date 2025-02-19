package api

import (
	"backend/errors"
	"backend/model"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewMovieController(movie_service *services.MovieService) *Movie_controller {
	return &Movie_controller{
		movie_service: movie_service,
	}
}

type Movie_controller struct {
	movie_service *services.MovieService
}

func (mc *Movie_controller) InsertNewMovie(c *gin.Context) {
	var movie_data model.Movie

	if err := c.ShouldBindBodyWithJSON(&movie_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "required data are either missing or malformed."})
		return
	}

	id, err := mc.movie_service.InsertNewMovie(movie_data)

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
	movie, err := mc.movie_service.GetMovieById(c.Param("id"))

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
	movie, err := mc.movie_service.GetListAllMovies()

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
	err := mc.movie_service.DeleteMovieById(c.Param("id"))

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
	movies, err := mc.movie_service.GetMoviesByPublisherId(c.Param("id"))

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