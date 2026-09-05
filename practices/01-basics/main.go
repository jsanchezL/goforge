package main

import "fmt"

func main() {
	name := "Jorge"
	var age int16 = 42
	height := 1.74
	isLearningGo := true

	const course = "GoForge"

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Learning Go:", isLearningGo)
	fmt.Println("Course:", course)

	fmt.Printf("Name: %s | Age: %d | Height: %.2f | Learning Go: %t\n", name, age, height, isLearningGo)

	var language string = "Go"

	fmt.Printf("%T\n", language)

	var name2 string = "Jorge Luis"
	var age2 int16 = 2026 - 1984
	var city string = "Xalapa"
	var yearsOfExperience int16 = 2026 - 2006
	var isDeveloper bool = true

	fmt.Printf("I'm %s. \nI live in %s.\nI'm %d years old.\nI've almost %d years of experience.\nNow I'm trying learn %s language.\nI'm a %t developer.\n", name2, city, age2, yearsOfExperience, language, isDeveloper)
}
