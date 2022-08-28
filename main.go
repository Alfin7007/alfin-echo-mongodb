package main

import (
	"explore/mongodb/config"
)

func main() {
	client := config.InitMongoDB()
}
