package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type cat struct {
	ID     string  `json:"id"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Url    string  `json:"url"`
	Breeds []breed `json:"breeds"`
}

type breed struct {
	Weight       weight `json:"weight"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	Temperament  string `json:"temperament"`
	Origin       string `json:"origin"`
	CountryCodes string `json:"countryCodes"`
	LifeSpan     string `json:"lifeSpan"`
	WikipediaUrl string `json:"wikipediaUrl"`
}

type weight struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}

var cats = []cat{
	{
		ID:     "0XYvRd7oD",
		Url:    "https://cdn2.thecatapi.com/images/0XYvRd7oD.jpg",
		Width:  1204,
		Height: 1445,
		Breeds: []breed{
			{
				Weight:       weight{Imperial: "7 - 10", Metric: "3 - 5"},
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

func getCats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cats)
}

func getCatByID(c *gin.Context) {
	id := c.Param("id")

	for _, cat := range cats {
		if cat.ID == id {
			c.IndentedJSON(http.StatusOK, cat)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cat not found"})
}

func postCat(c *gin.Context) {
	var newCat cat

	if err := c.BindJSON(&newCat); err != nil {
		return
	}

	cats = append(cats, newCat)
	c.IndentedJSON(http.StatusCreated, newCat)
}

func main() {
	router := gin.Default()

	router.GET("/cats", getCats)
	router.GET("/cats/:id", getCatByID)
	router.POST("/cats", postCat)

	router.Run("localhost:8080")
}
