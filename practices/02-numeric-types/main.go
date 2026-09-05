package main

import "fmt"

func main() {
	var product string = "Tempered Glass 10mm"
	var quantity int8 = 3
	var unitPrice float32 = 1250.50
	var discount float32 = 10.00
	var tax float32 = 16.00

	var subtotal float32 = float32(quantity) * unitPrice
	var discountAmount float32 = subtotal * discount / 100
	var subtotalAfterDiscount float32 = subtotal - discountAmount
	var taxAmount float32 = subtotalAfterDiscount * tax / 100
	var finalTotal float32 = subtotalAfterDiscount + taxAmount

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
