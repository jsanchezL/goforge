package main

import "fmt"

func main() {
	age := 20

	if age >= 18 {
		fmt.Println("You're adult")
	} else {
		fmt.Println("You're minor")
	}

	score := 85

	if score >= 90 {
		fmt.Println("Excellent")
	} else if score >= 80 {
		fmt.Println("Very good")
	} else if score >= 70 {
		fmt.Println("Good")
	} else {
		fmt.Println("Needs improvement")
	}

	hasLicense := true

	if age >= 18 && hasLicense {
		fmt.Println("Can drive")
	}

	isAdmin := false
	isOwner := true

	if isAdmin || isOwner {
		fmt.Println("Access granted")
	}

	isBlocked := false

	if !isBlocked {
		fmt.Println("Use can continue")
	}

	product := "Tempered Glass 10mm"
	subtotal := 3751.50
	discount := 0

	if subtotal >= 5000 {
		discount = 15
	} else if subtotal >= 3000 {
		discount = 10
	}

	discountAmount := subtotal * float64(discount) / 100
	finalSubtotal := subtotal - discountAmount
	fmt.Printf("Product: %s.\nSubtotal: $%.2f\nDiscount applied: %d%% \nDiscount amount: $%.2f\nFinal subtotal: %.2f", product, subtotal, discount, discountAmount, finalSubtotal)
}
