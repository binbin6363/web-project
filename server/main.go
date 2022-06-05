package main

import (
	"flag"
	"log"

	"github.com/binbin6363/web-project/server/config"
	"github.com/binbin6363/web-project/server/middleware"
	"github.com/binbin6363/web-project/server/route"
	"github.com/gin-gonic/gin"
)

var configPath = flag.String("config", "", "Input Your config")

func main() {

	//flag.StringVar(configPath, "config", "", "Input Your configo")
	flag.Parse()

	log.Print(*configPath)
	config.Init(*configPath)
	engine := gin.Default()
	engine.Use(middleware.Cors())

	if err := middleware.NewMySQL(config.GetConfig().Db.Addr); err != nil {
		log.Fatal("init db failed")
	}

	go middleware.NewTick().Tick()

	route.Route(engine)

	if err := engine.Run(config.GetConfig().Server.Addr); err != nil {
		log.Fatal("start failed")
	}
}
