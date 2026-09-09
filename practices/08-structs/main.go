package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Stock    int
	Category string
}

func main() {
	products := []Product{
		{
			Name:     "Tempered Glass 10mm",
			Price:    1250.50,
			Stock:    8,
			Category: "Glass",
		},
		{
			Name:     "Aluminium",
			Price:    980.0,
			Stock:    18,
			Category: "Metal",
		},
		{
			Name:     "Wood",
			Price:    2100.75,
			Stock:    4,
			Category: "Wood",
		},
		{
			Name:     "Silicone",
			Price:    150.79,
			Stock:    5,
			Category: "Consumable",
		},
	}

	inventoryValue := 0.0
	mostExpensiveProduct := products[0]
	lowStockProductCount := 0
	aluminiumIndex := -1

	for i, product := range products {
		fmt.Printf(
			"%s | $%.2f | Stock: %d | %s\n",
			product.Name,
			product.Price,
			product.Stock,
			product.Category,
		)
		inventoryValue += (product.Price * float64(product.Stock))

		if product.Price > mostExpensiveProduct.Price {
			mostExpensiveProduct = product
		}

		if product.Stock < 10 {
			lowStockProductCount++
		}

		if product.Name == "Aluminium" {
			products[i].Stock += 5
			aluminiumIndex = i
		}
	}

	fmt.Println()
	fmt.Printf("Inventory value: $%.2f\n", inventoryValue)
	fmt.Println("Most expensive:", mostExpensiveProduct.Name)
	fmt.Println("Low-stock products:", lowStockProductCount)
	if aluminiumIndex >= 0 {
		fmt.Println("Updated Aluminium stock:", products[aluminiumIndex].Stock)
	}
}
