package main

import (
	"GoBlog/model"
	"GoBlog/routes"
)

func main() {
	// test for dev
	model.InitDb()
	routes.InitRouter()

}
