package main

import (
	"github.com/bootcamp-go/desafio-cierre-testing/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	r.Run(":8080")
}
