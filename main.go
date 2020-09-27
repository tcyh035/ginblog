package main

import (
	"ginblog/model"
	"ginblog/routes"
	"ginblog/utils"
)

func main() {
	model.InitDb()

	engine := routes.InitRouter()

	engine.Run(utils.HTTPPort)
}
