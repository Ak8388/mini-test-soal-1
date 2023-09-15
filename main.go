package main

import "fmt"

func main() {
	var value []int

	for i := 1; i <= 100; i++ {
		value = append(value, i)
	}

	for number := len(value) - 1; number >= 0; number-- {

		if PrimaValidate(value[number]) == true {
			fmt.Print()
		} else {
			if value[number]%3 == 0 && value[number]%5 == 0 {
				fmt.Print(" FooBar,")
			} else if value[number]%5 == 0 {
				fmt.Print(" Bar,")
			} else if value[number]%3 == 0 {
				fmt.Print(" Foo,")
			} else {
				fmt.Printf(" %d,", value[number])
			}
		}
	}
}

func PrimaValidate(number int) bool {
	validate := 0

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			validate++
		}
	}

	if validate == 2 && number > 1 {
		return true
	} else {
		return false
	}

}
