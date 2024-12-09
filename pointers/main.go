package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {

	name := "tifa"

	// updateName(name) for updateName(x string) format function

	fmt.Println("memory address of name is: ", &name)

	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)

	fmt.Println("original name =", name)
	updateName(m)
	fmt.Println("updated name =", name)

}

/*
|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
|  "tifa" |  p0x001 |
|---------|---------|
*/
