package server

import (
	"receipt/handler"

	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	router := gin.Default()
	router.POST("/receipts/process", handler.ProcessReceipts)
	router.GET("/receipts/:id/points", handler.GetPoints)
	return router
}
