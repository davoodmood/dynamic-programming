package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var column, row uint = 18, 18

// we use md5 hashing to covert two integer valiables m & n as a unique key with type String.
var memory map[string]uint = make(map[string]uint)

var print = message.NewPrinter(language.English)

func main() {
	possibilities := gridTraveler(column, row, memory)
	// using text module Printf instead of fmt to display the return digit with comma(,) for thousands.
	print.Printf("the GIAGANTIC (%d,%d) grid has | %d | number of possibilities.\n", column, row, possibilities)
	print.Printf("the Space is: %v\n", len(memory))
	var i uint
	for i = 0; i <= 10; i++ {
		column = 2 + i
		row = 3 + i
		possibilities = gridTraveler(column, row, memory)
		print.Printf("the (%d,%d) grid has | %d | number of possibilities.\n", column, row, possibilities)
	}
}

func gridTraveler(m uint, n uint, memory map[string]uint) uint {
	key := fmt.Sprintf("f(%d,%d)", m, n)

	if value, ok := memory[key]; ok {
		return value
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	memory[key] = gridTraveler(m-1, n, memory) + gridTraveler(m, n-1, memory)
	return memory[key]
}
