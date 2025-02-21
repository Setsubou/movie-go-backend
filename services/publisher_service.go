package services

import (
	"backend/errors"
	"backend/model"
	"backend/repository"
)

type PublisherService struct {
	repository repository.PublisherRepository
}

func (p *PublisherService) GetListAllPublishers() ([]*model.Publisher, error) {
	publishers, err := p.repository.GetListAllPublishersName()
	var converted_publishers []*model.Publisher

	if err != nil {
		return nil, errors.ErrNotFound.SetMessage("no publisher found")
	}

	for _, publishers := range *publishers {
		publisher := model.Publisher{
			Id:             publishers.ID.String(),
			Publisher_name: publishers.PublisherName,
		}

		converted_publishers = append(converted_publishers, &publisher)
	}

	return converted_publishers, nil
}

func NewPublisherService(repository repository.PublisherRepository) *PublisherService {
	return &PublisherService{
		repository: repository,
	}
}
