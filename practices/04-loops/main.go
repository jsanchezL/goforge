package main

import "fmt"

func main() {
	count := 0

	for count < 5 {
		fmt.Println(count)
		count++
	}

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	for n := 1; n <= 20; n++ {
		if n == 18 {
			break
		}

		if n == 13 {
			continue
		}

		if n%2 == 0 {
			fmt.Printf("%d .- Even\n", n)
		} else {
			fmt.Printf("%d .- Odd\n", n)
		}
	}
}
