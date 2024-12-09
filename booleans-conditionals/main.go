package main

import "fmt"

func main() {

	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is greater than or equal to 40")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Printf("continuing at pos %v \n", index)
			continue
		}

		if index > 2 {
			fmt.Printf("breaking at pos %v \n", index)
			break
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

}
