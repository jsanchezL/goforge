package main

import "fmt"

func main() {
	product := "Tempered Glass 10mm"
	quantity := 3
	unitPrice := 1250.50
	discount := 10.00
	tax := 16.00

	subtotal := float64(quantity) * unitPrice
	discountAmount := subtotal * discount / 100
	subtotalAfterDiscount := subtotal - discountAmount
	taxAmount := subtotalAfterDiscount * tax / 100
	finalTotal := subtotalAfterDiscount + taxAmount

	fmt.Println("Product: ", product)
	fmt.Println("Quantity: ", quantity)
	fmt.Printf("Unit price: $%.2f\n", unitPrice)
	fmt.Println("")
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("Discount: $%.2f\n", discountAmount)
	fmt.Printf("After discount: $%.2f\n", subtotalAfterDiscount)
	fmt.Printf("Tax: $%.2f\n", taxAmount)
	fmt.Printf("Final total: $%.2f\n", finalTotal)
}
