package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// changing the nth index of the Fibonachi value here to get it's value
var nth uint = 1000

func main() {
	memory := make(map[uint]uint)
	fibNum := fib(nth, memory)
	print := message.NewPrinter(language.English)
	print.Printf("the %dth value fibonachi number is: %d \n", nth, fibNum)
	for nth = 0; nth <= 10; nth++ {
		fibNum = fib(nth, memory)
		print.Printf("the %dth value fibonachi number is: %d \n", nth, fibNum)
	}
}

func fib(n uint, memory map[uint]uint) uint {
	if value, ok := memory[n]; ok {
		return value
	}
	if n <= 2 {
		return 1
	}
	memory[n] = fib(n-1, memory) + fib(n-2, memory)
	return memory[n]
}
