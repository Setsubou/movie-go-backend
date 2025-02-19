package api

import (
	"backend/errors"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewPublisherController(publisher_service *services.PublisherService) *Publisher_controller {
	return &Publisher_controller{
		publisher_service: publisher_service,
	}
}

type Publisher_controller struct {
	publisher_service *services.PublisherService
}

func (pc *Publisher_controller) GetListAllPublishersName(c *gin.Context) {
	publisher, err := pc.publisher_service.GetListAllPublishers()

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
