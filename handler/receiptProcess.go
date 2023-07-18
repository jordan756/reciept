package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"receipt/structs"
	utils "receipt/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var M map[string]int
var validate *validator.Validate

func init() {
	M = map[string]int{}
	validate = validator.New()
	validate.RegisterValidation("time", func(fl validator.FieldLevel) bool {
		_, err := time.Parse("15:04", fl.Field().String())
		return err == nil
	})
	validate.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		_, err := time.Parse("2006-01-02", fl.Field().String())
		return err == nil
	})
	validate.RegisterValidation("money", func(fl validator.FieldLevel) bool {
		amount, err := strconv.ParseFloat(fl.Field().String(), 64)
		if err != nil {
			return false
		}
		if amount < 0 {
			return false
		}
		regex := regexp.MustCompile(`^[0-9]+(\.[0-9]+)$`)
		return regex.MatchString(fmt.Sprintf("%.2f", amount))
	})
}

func ProcessReceipts(c *gin.Context) {
	var receipt structs.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := validate.Struct(receipt)
	if err != nil {
		fmt.Println("Validation error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uuid.New()
	value := utils.CaculatePoints(receipt)
	M[id.String()] = value
	response := structs.IdResponse{Id: id.String()}
	c.IndentedJSON(http.StatusOK, response)
	//c.IndentedJSON(http.StatusOK, M)
}

func GetPoints(c *gin.Context) {

	id := c.Params.ByName("id")
	points, ok := M[id]
	
	if !ok {

		c.IndentedJSON(http.StatusBadRequest, "invalid id")
		return
	}

	p := structs.PointsResponse{Points: points}
	c.IndentedJSON(http.StatusOK, p)
}
