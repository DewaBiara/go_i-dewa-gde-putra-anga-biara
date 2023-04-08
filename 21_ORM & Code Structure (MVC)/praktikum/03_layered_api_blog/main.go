package main

import (
	"03_layered_api_blog/config"
	"03_layered_api_blog/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
