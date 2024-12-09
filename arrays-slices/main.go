package main

import "fmt"

func main() {

	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 35, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(ages)

	fmt.Println(ages, "ages size =", len(ages))
	fmt.Println(names, "names size =", len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60} // gets initialized as slice

	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores)

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}
