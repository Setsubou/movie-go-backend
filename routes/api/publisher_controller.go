package api

import (
	"backend/errors"
	"backend/repository"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Publisher_controller struct {
	repository repository.PublisherRepository
}

func (pc *Publisher_controller) GetListAllPublishersName(c *gin.Context) {
	publisher, err := services.NewPublisherService(pc.repository).GetListAllPublishers()

	if err != nil {
		if internalErr, ok := err.(*errors.InternalError); ok {
			c.JSON(internalErr.Code, gin.H{"error": internalErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, publisher)
}

func NewPublisherController(repository repository.PublisherRepository) *Publisher_controller {
	return &Publisher_controller{
		repository: repository,
	}
}
