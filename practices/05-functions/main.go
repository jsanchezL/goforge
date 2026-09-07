package main

import "fmt"

func calculateSubtotal(quantity int, unitPrice float64) float64 {
	return float64(quantity) * unitPrice
}

func calculateDiscount(subtotal float64, discount float64) float64 {
	return subtotal * discount / 100
}

func calculateTax(amount float64, taxRate float64) float64 {
	return amount * taxRate / 100
}

func getDiscountRate(subtotal float64) float64 {
	discount := 0.0

	if subtotal >= 5000 {
		discount = 15.0
	} else if subtotal >= 3000 {
		discount = 10.0
	}

	return discount
}

func main() {
	quantity := 3
	unitPrice := 1250.50
	taxRate := 16.00
	product := "Tempered Glass 10mm"

	subtotal := calculateSubtotal(quantity, unitPrice)
	discountRate := getDiscountRate(subtotal)
	discountAmount := calculateDiscount(subtotal, discountRate)
	subtotalAfterDiscount := subtotal - discountAmount
	taxAmount := calculateTax(subtotalAfterDiscount, taxRate)
	finalTotal := subtotalAfterDiscount + taxAmount

	fmt.Println("Product:", product)
	fmt.Println("Quantity:", quantity)
	fmt.Printf("Unit price: $%.2f\n", unitPrice)
	fmt.Println()
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("Discount rate: %.0f%%\n", discountRate)
	fmt.Printf("Discount amount: $%.2f\n", discountAmount)
	fmt.Printf("After discount: $%.2f\n", subtotalAfterDiscount)
	fmt.Printf("Tax rate: %.0f%%\n", taxRate)
	fmt.Printf("Tax amount: $%.2f\n", taxAmount)
	fmt.Printf("Final total: $%.2f\n", finalTotal)
}
