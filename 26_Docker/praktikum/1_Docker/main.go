package main

import (
	"1_JWT_Auth/config"
	"1_JWT_Auth/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":80"))
}
