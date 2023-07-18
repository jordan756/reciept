package utils

import (
	"receipt/structs"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCaculatePointsSuccess(t *testing.T) {
	recipt := structs.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []structs.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}
	sum := CaculatePoints(recipt)
	assert.Equal(t, 28, sum)
}

func TestCaculatePointsFailure(t *testing.T) {
	recipt := structs.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01111",
		PurchaseTime: "13:01111",
		Total:        "35.35",
		Items: []structs.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}
	sum := CaculatePoints(recipt)
	assert.Equal(t, -1, sum)
}