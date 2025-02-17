package model

import db "backend/db/sqlc"

type Country struct {
	Id           string `json:"id"`
	Country_name string `json:"country_name"`
}

func ConvertCountryFromRepository(c db.Country) (*Country, error) {
	return &Country{
		Id: c.ID.String(),
		Country_name: c.CountryName,
	}, nil
}