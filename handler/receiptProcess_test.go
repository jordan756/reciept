package handler

import (
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"receipt/structs"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestProcessReceiptsSuccess(t *testing.T) {
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
	jsonBody, err := json.Marshal(recipt)
	if err != nil {
		t.Fatal(err)
	}
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()

	router := gin.Default()

	router.POST("/receipts/process", ProcessReceipts)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}
func TestProcessReceiptsFailure(t *testing.T) {
	recipt := structs.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-91",
		PurchaseTime: "13:01:9999",
		Total:        "35.35",
		Items: []structs.Item{

			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}
	jsonBody, err := json.Marshal(recipt)
	if err != nil {
		t.Fatal(err)
	}
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()

	router := gin.Default()

	router.POST("/receipts/process", ProcessReceipts)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGetPointsSuccess(t * testing.T) {
	recipt := structs.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []structs.Item{	
		},
	}
	jsonBody, err := json.Marshal(recipt)
	if err != nil {
		t.Fatal(err)
	}
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()

	router := gin.Default()

	router.POST("/receipts/process", ProcessReceipts)
	router.GET("/receipts/:id/points", GetPoints)

	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}


	var response structs.IdResponse
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		t.Fatal(err)
	}
	
	
	req, _ = http.NewRequest("GET","/receipts/"+response.Id+"/points",nil)
	w = httptest.NewRecorder()

	
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}
func TestGetPointsFailure(t * testing.T) {
	req, _ := http.NewRequest("GET", "/receipts/fakeId1234/points",nil)
	w := httptest.NewRecorder()

	router := gin.Default()

	router.GET("/receipts/:id/points", GetPoints)
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}
}
