package api

import (
	"backend/errors"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGenreController(genre_service *services.GenreService) *Genre_controller {
	return &Genre_controller{
		genre_service: genre_service,
	}
}

type Genre_controller struct {
	genre_service *services.GenreService
}

func (gc *Genre_controller) GetAllGenres(c *gin.Context) {
	genres, err := gc.genre_service.GetAllGenres()

	if err != nil {
		if internalErr, ok := err.(*errors.InternalError); ok {
			c.JSON(internalErr.Code, gin.H{"error": internalErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, genres)
}
