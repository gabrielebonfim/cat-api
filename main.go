package main

import (
	"github.com/gabrielebonfim/cat-api/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run("localhost:8080")
}
