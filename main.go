package main

import (
	"reciept/server"
)

func main() {

	server := server.CreateServer()
	server.Run("localhost:8080")

}
