package main

import "fmt"

func main() {
	productPrices := map[string]float64{
		"Tempered Glass 10mm": 1250.50,
		"Aluminium":           980.00,
		"Wood":                2100.75,
	}

	fmt.Println(productPrices)

	inventory := map[string]int{
		"Tempered Galss": 8,
		"Aluminium":      15,
		"Wood":           4,
		"Hardware":       20,
	}

	inventory["Silicone"] = 12
	quantity, exists := inventory["Aluminium"]

	if exists {
		inventory["Aluminium"] = quantity + 3
	}

	q, exists := inventory["Wood"]

	if exists {
		fmt.Printf("We have %d units of Wood in inventory\n", q)
	}

	_, exists = inventory["Steel"]

	if !exists {
		fmt.Println("Steel does not exist in inventory")
	}

	delete(inventory, "Hardware")

	totalUnits := 0
	lowStockCount := 0
	for product, quantity := range inventory {
		fmt.Printf("%s: %d\n", product, quantity)
		totalUnits += quantity

		if quantity < 10 {
			lowStockCount++
		}
	}
	fmt.Println("Total units:", totalUnits)
	fmt.Println("Count fewer 10 units:", lowStockCount)
}
