package handlers

import (
	"net/http"

	"github.com/gabrielebonfim/cat-api/models"
	"github.com/gabrielebonfim/cat-api/repository"
	"github.com/gin-gonic/gin"
)

func GetCats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repository.FindCats())
}

func GetCatByID(c *gin.Context) {
	id := c.Param("id")
	cat := repository.FindCatById(id)

	if cat.ID != "" {
		c.IndentedJSON(http.StatusOK, cat)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cat not found"})
}

func PostCat(c *gin.Context) {
	var newCat models.Cat

	if err := c.BindJSON(&newCat); err != nil {
		return
	}

	repository.InsertCat(newCat)
	c.IndentedJSON(http.StatusCreated, newCat)
}
