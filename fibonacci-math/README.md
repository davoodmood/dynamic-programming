# fib(n)

## The problem

Write a function `fib(n)` that takes in a nuber as an argument. The function should return the n-th number of the Fibonacci sequence.

## Tips

The 1st and 2nd number of the sequence is 1.
Togenerate the next number of the sequence, we sum the previous two.

`fib(n): 1, 1, 2, 3, 5, 8, 13, 21, 34, ...`

## Issue resolved

Eventhough when Go programming language is sublime in handling the processing CPU intensive tasks. There is a problem when you try to calculate the Fibonacci value for n-th number when `n >= 50`. So the code uses `memorization` strategy in Dynamic Programming with an algorythm to speed up the calculation from a minute long process (with GO! more with other languages, like Python or Node.js) to reduce it to miliseconds.

# The program

## Specify the n-th Fibonacci

You can change `line 9` to specify what is the n-th number you want the Fibonacci number.

``` Go
// changing the nth index of the Fibonachi value here to get it's value
var nth uint = 1000
```

## Run the program

Move to the folder where the `fibMath.go` is located on you command line or bash.
Type `Go run .` to run the program and get the response. I have also calculated the sequence from 1st to 10th, you will see the responses on your console.



