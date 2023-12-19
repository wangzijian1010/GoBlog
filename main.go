package main

import (
	"GoBlog/model"
	"GoBlog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()

}
