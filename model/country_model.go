package model

import db "backend/db/sqlc"

type Country struct {
	Id           string `json:"id,omitempty"`
	Country_name string `json:"country_name,omitempty"`
}

func ConvertCountryFromDTO(c db.Country) Country {
	return Country{
		Id:           c.ID.String(),
		Country_name: c.CountryName,
	}
}
