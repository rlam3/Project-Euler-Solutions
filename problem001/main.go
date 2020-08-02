// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func divisibleBy3or5(value int) bool {

	if value%3 == 0 || value%5 == 0 {

		return true

	}

	return false

}

func main() {

	var allDivisible []int

	for i := 1; i < 1000; i++ {
		if divisibleBy3or5(i) {
			allDivisible = append(allDivisible, i)
		}
	}

	fmt.Println(allDivisible)

	var summation int

	for _, b := range allDivisible {
		summation += b
	}

	fmt.Println(summation)
	fmt.Println("Hello, World")
	// Loop through allDivisible and sum up

}
