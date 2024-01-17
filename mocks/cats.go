package mocks

import (
	"github.com/gabrielebonfim/cat-api/models"
)

var Cats = []models.Cat{
	{
		ID:     "0XYvRd7oD",
		Url:    "https://cdn2.thecatapi.com/images/0XYvRd7oD.jpg",
		Width:  1204,
		Height: 1445,
		Breeds: []models.Breed{
			{
				Weight:       models.Weight{Imperial: "7 - 10", Metric: "3 - 5"},
				Id:           "abys",
				Name:         "Abyssinian",
				Temperament:  "Active, Energetic, Independent, Intelligent, Gentle",
				Origin:       "Egypt",
				CountryCodes: "EG",
				LifeSpan:     "14 - 15",
				WikipediaUrl: "https://en.wikipedia.org/wiki/Abyssinian_(cat)",
			},
		},
	},
}
