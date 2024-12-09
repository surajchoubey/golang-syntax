package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}

	// add tip
	fs += fmt.Sprintf("%25v ...$%v\n", "tip:", b.tip)
	total += b.tip

	// total
	fs += fmt.Sprintf("%25v ...$%0.2f", "total:", total)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip // no need to do (*b).tip = tip, Go automatically resolves that
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	/*
		when we're saving data to a file, we need to save the file in byte slice format
		or slice of bytes format
		so we have to convert them in byte slice format
		then we can save it into a file
	*/
	data := []byte(b.format())
	filename := b.name + ".txt"
	err := os.WriteFile("bills/"+filename, data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("bill was saved to bills/" + filename)
}
