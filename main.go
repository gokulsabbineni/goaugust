package main

import "fmt"

func main() {
	numbers := []int{10,20}
	fmt.Println(len(numbers), cap(numbers))

	numbers = append(numbers, 30)
	fmt.Println(len(numbers), cap(numbers))
	fmt.Println(numbers)

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	ages := map[string]int {
		"john": 28,
		"go": 28, 
	}

	age, exists := ages["johnn"]

	if exists {
		fmt.Println(age)
	} else {
		fmt.Println("not present")
	}

	for key, value := range ages {
		fmt.Println(key, value)
	}
}