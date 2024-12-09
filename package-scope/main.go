package main

import "fmt"

var score = 99.5 // showScore() CAN USE this variable as it is within package scope

func main() {

	// var score = 99.5 : ERROR

	sayHello("mario")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()

}

// to run this use:
// go run main.go greetings.go
