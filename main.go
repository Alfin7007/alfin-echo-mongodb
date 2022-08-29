package main

import (
	"explore/mongodb/config"
	"explore/mongodb/factory"
	"explore/mongodb/middlewares"
	"explore/mongodb/routes"
)

func main() {
	client := config.InitMongoDB()
	presenter := factory.InitFactory(client)
	e := routes.InitRoute(presenter)

	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
