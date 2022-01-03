package main

import (
	"github.com/binbin6363/web-project/server/route"
	"github.com/binbin6363/web-project/server/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.Cors())

	route.Route(engine)

	engine.Run(":8080")
}