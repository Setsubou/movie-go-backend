package api

import (
	"backend/errors"
	"backend/repository"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Genre_controller struct {
	repository repository.GenreRepository
}

func (gc *Genre_controller) GetAllGenres(c *gin.Context) {
	genres, err := services.NewGenreService(gc.repository).GetAllGenres()

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

func NewGenreController(repository repository.GenreRepository) *Genre_controller {
	return &Genre_controller{
		repository: repository,
	}
}
