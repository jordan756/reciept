package utils

import (
	"math"
	"receipt/structs"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func CaculatePoints(receipt structs.Receipt) (int,error) {
	sum := 0
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char)  || unicode.IsDigit(char) {
			sum += 1
		}
	}
	//fmt.Println("retailer added ", sum, " points") // caculate retailer

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return sum,err
	}

	if math.Round(total) == total { //check if no cents remainder in total
		sum += 50
	//	fmt.Println("total added ", 50, " points")
	}

	if math.Mod(total, .25) == 0 { //check if total is multiple .25
		sum += 25
	//	fmt.Println("total added ", 25, " points")
	}

	quotient := len(receipt.Items) / 2 //give points by item list length 2:5
	sum += (quotient * 5)

	//fmt.Println("Items length added ", quotient*5, " points")

	for _, item := range receipt.Items {
		description := strings.TrimSpace(item.ShortDescription) //give points if descrip is / by 3
		if math.Mod(float64(len(description)), 3) == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			sum += int(math.Ceil(.2 * price))
	//		fmt.Println(description, ": added ", math.Ceil(.2*price), " points")
		}
	}

	layout := "2006-01-02"
	date, err := time.Parse(layout, receipt.PurchaseDate)
	if err != nil {
	//	fmt.Println("Error:", err)
		return sum,err
	}
	if (date.Day() % 2) == 1 { // if odd day, add 6 points
		sum += 6
	//	fmt.Println(date, " added ", 6, " points")
	}

	layout = "15:04"
	purchaseTime, err := time.Parse(layout, receipt.PurchaseTime) // add 10 points if between 2-4
	if err != nil {
	//	fmt.Println("Error:", err)
		return sum,err
	}
	two, _ := time.Parse(layout, "14:00")
	four, _ := time.Parse(layout, "16:00")
	if purchaseTime.After(two) && purchaseTime.Before(four) {
		sum += 10
	//	fmt.Println(purchaseTime, " added ", 10, " points")
	}

	return sum,nil

}
