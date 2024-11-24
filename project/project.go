package project

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category int    `json:"category"`
}

func GetAllProducts() {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	var products []Product
	if err := json.Unmarshal(bodyBytes, &products); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Products:")
	for _, product := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %d, Category: %d\n",
			product.Id, product.Name, product.Price, product.Category)
	}
}
