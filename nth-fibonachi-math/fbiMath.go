package main

import (
	"fmt"
)

// changing the nth index of the Fibonachi value here to get it's value
var nth uint64 = 50

func main() {
	fibNum := fib(nth)
	fmt.Printf("the %vth value fibonachi number is: %v \n", nth, fibNum)
	// for nth = 7; nth < 11; nth++ {
	// 	fibNum = fib(nth)
	// 	fmt.Printf("the %vth value fibonachi number is: %v\n", nth, fibNum)
	// }
}

func fib(n uint64) uint64 {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
