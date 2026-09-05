package main

import "fmt"

func main() {
	var product string = "Tempered Glass 10mm"
	var qyt int8 = 3
	var unit_price float32 = 1250.50
	var discount float32 = 10.00
	var tax float32 = 16.00

	var subtotal float32 = float32(qyt) * unit_price
	var discount_amount float32 = subtotal / discount
	var subtotal_after_discount float32 = subtotal - discount_amount
	var tax_amount float32 = (subtotal_after_discount * tax) / 100
	var final_total float32 = subtotal_after_discount + tax_amount

	fmt.Println("Product: ", product)
	fmt.Println("Quantity: ", qyt)
	fmt.Printf("Unit price: $%.2f\n", unit_price)
	fmt.Println("")
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("Discount: $%.2f\n", discount_amount)
	fmt.Printf("After discount: $%.2f\n", subtotal_after_discount)
	fmt.Printf("Taxes: $%.2f\n", tax_amount)
	fmt.Printf("Final total: $%.2f\n", final_total)
}
