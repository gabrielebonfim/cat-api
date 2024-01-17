package repository

import (
	"github.com/gabrielebonfim/cat-api/mocks"
	"github.com/gabrielebonfim/cat-api/models"
)

func FindCats() []models.Cat {
	return mocks.Cats
}

func FindCatById(id string) models.Cat {
	for _, cat := range mocks.Cats {
		if cat.ID == id {
			return cat
		}
	}
	return models.Cat{}
}

func InsertCat(newCat models.Cat) {
	mocks.Cats = append(mocks.Cats, newCat)
}
