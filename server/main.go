package main

import (
	"./route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	route.Route(r)

	r.Run(":8080")
}