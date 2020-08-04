// https://projecteuler.net/problem=3

// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

// What do we know about prime numbers?
// they are non divisible by any other number but itself.
// a composite number

// Approach 1
// make a list of prime numbers and test out each prime number until you reach the highest one.
// Highest prime number should not be larger than the number given.
// Make a list of prime numbers that the current number given is a prime number of

// Approach 1.1
// Use the factored down value to continue solving instead of every number??

// Approach 2
// split the number as a b-tree? and
// use prime factors to break it down until it cannot be broken down with prime factors
// Components of a Tree:
// Node
//		value
//		is value of node prime?
//		children -- defines left and right

package main

import "fmt"

func isPrime(value int) bool {
	var isValuePrime = false

	fmt.Printf("Checking....%d isPrime...\n", value)
	for i := 2; i <= value; i++ {

		if value%i == 0 && value%2 == 0 {
			isValuePrime = false
		}

		// If there is a number divisible under value, it is not a prime number
		if value%i == 0 {
			fmt.Printf("%d is NOT PRIME\n", value)
			isValuePrime = false
		}

		if value%i == 0 && value == i {
			fmt.Printf("%d is PRIME\n", value)
			isValuePrime = true
		}
	}
	return isValuePrime
}

func main() {
	fmt.Println("problem 3")
	// fmt.Printf("%t\n", isPrime(3))
	var listOfPrimesDivisible []int
	const largeNumber = 111

	for i := 2; i <= largeNumber; i++ {
		if isPrime(i) {
			if largeNumber%i == 0 {
				// fmt.Printf("Checking...%d\n", i)
				listOfPrimesDivisible = append(listOfPrimesDivisible, i)
			}
		}
	}

	fmt.Println(listOfPrimesDivisible)
}
