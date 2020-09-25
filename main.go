package main

import (
	"ginblog/routes"
	"ginblog/utils"
)

func main() {
	engine := routes.InitRouter()

	engine.Run(utils.HTTPPort)
}
