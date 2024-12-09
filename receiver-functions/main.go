package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")

	fmt.Println(mybill)

	fmt.Println(mybill.format())

}

/*
f(x)

// val
y = f(x)

// ref
f(x)

f (classname object)
*/
