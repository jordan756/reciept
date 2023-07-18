package handler

import (
	"fmt"
	"net/http"

	"reciept/structs"
	utils "reciept/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var M map[string]int

func init() {
	M = map[string]int{}
}

func ProcessReceipts(c *gin.Context) {
	var reciept structs.Reciept

	if err := c.ShouldBindJSON(&reciept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uuid.New()
	value := utils.CaculatePoints(reciept)
	M[id.String()] = value
	response := structs.IdResponse{Id: id.String()}
	c.IndentedJSON(http.StatusOK, response)
}

func GetPoints(c *gin.Context) {

	id := c.Params.ByName("id")
	fmt.Println("id")
	points, ok := M[id]
	if ok {
		p := structs.PointsResponse{Points: points}
		c.IndentedJSON(http.StatusOK, p)
		return
	} else {
		points = 12222
	}
	c.IndentedJSON(http.StatusOK, M)
}
