package main

import "fmt"

func main() {
	var prices [3]float64
	prices[0] = 1250.50
	prices[1] = 980.00
	prices[2] = 2100.75
	fmt.Println(prices)

	prices1 := [3]float64{1250.50, 980.00, 2100.75}
	fmt.Println(prices1)

	prices2 := []float64{1250.50, 980.00, 2100.75}
	fmt.Println(prices2)

	prices2 = append(prices2, 750.25)
	fmt.Println(len(prices2))

	for index, price := range prices {
		fmt.Println(index, price)
	}

	for _, price := range prices2 {
		fmt.Println(price)
	}

	prices3 := []float64{1250.50, 980.00, 2100.75, 450.25, 3200.00}
	total := 0.0
	highest := prices3[0]
	moreThan1000Count := 0

	prices3 = append(prices3, 1750.25)

	for i, p := range prices3 {
		fmt.Printf("Price %d: $%.2f\n", i+1, p)
		total += p

		if p > highest {
			highest = p
		}

		if p > 1000 {
			moreThan1000Count++
		}
	}
	fmt.Printf("Total: $%.2f\n", total)
	fmt.Printf("Average: $%.2f\n", (total / float64(len(prices3))))
	fmt.Printf("Highest price: $%.2f\n", highest)
	fmt.Println("Prices over $1000:", moreThan1000Count)
}
