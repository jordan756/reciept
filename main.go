package main

import (
	"log"
	"receipt/server"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("local_app.env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error reading .env file:", err)
	}
	server := server.CreateServer()
	//server.Run("localhost:8080")
	server.Run(viper.GetString("host"))
}
