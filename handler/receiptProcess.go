package handler

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"

	"receipt/config"
	"receipt/structs"
	utils "receipt/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var M map[string]int 
var allReciepts map[[32]byte]bool  //prevent duplicates of same reciepts
var validate *validator.Validate
var rw   sync.RWMutex

func init() {
	M = map[string]int{}
	allReciepts = map[[32]byte]bool{} 
	validate = validator.New()
	//add custom validations for time, date and money.
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
		regex := regexp.MustCompile(`[0-9]*\.[0-9]{2}$`)
		return regex.MatchString(fl.Field().String())
	})
}

func ProcessReceipts(c *gin.Context) {
	var receipt structs.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{Error: config.JsonError + err.Error()})
		return
	}
	err := validate.Struct(receipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{Error: config.ValidationError + err.Error()})
		return
	}
	text, err := json.Marshal(receipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{Error: config.JsonError + err.Error()})
		return
	}
	hash := sha256.Sum256([]byte(text))

	rw.Lock() //lock before accessing map
	defer rw.Unlock()


	if _,ok := allReciepts[hash]; ok {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{Error: config.RepeatError})
		return
	}
	allReciepts[hash] = false

	id := uuid.New()
	value, err := utils.CaculatePoints(receipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{Error: config.CaculatePointsFailed + err.Error()})
		return
	}
	M[id.String()] = value


	response := structs.IdResponse{Id: id.String()}
	c.IndentedJSON(http.StatusOK, response)
	//c.IndentedJSON(http.StatusOK, M)
}

func GetPoints(c *gin.Context) {
	rw.RLock()
    defer rw.RUnlock()
	id := c.Params.ByName("id")
	points, ok := M[id]
	
	if !ok {

		c.IndentedJSON(http.StatusBadRequest, structs.ErrorResponse{Error: config.GetIdError})
		return
	}

	p := structs.PointsResponse{Points: points}
	c.IndentedJSON(http.StatusOK, p)
}
