package repository

import db "backend/db/sqlc"

type PublisherRepository interface {
	GetListAllPublishersName() (*[]db.GetAllPublishersNameRow, error)
}
