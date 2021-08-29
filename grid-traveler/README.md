# Grid Traveler

## The problem

Say that you are a traveler on a two dimentional grid. You begin in the top-left corner and your goal is to travel to the bottom-right corner. You may only move down or right.

In how many ways can you travel to the bottom-right on a grid with dimentions `Columns * Rows`

## Tips

Write a function `gridTraveler(m, n)` to calculate this.
Togenerate the next number of the sequence, we sum the previous two.

e.g. `gridTraveler(2,3) => 3`

| START |       |       |
|       |       |  END  |

## Issue resolved

Eventhough when Go programming language is sublime in handling the processing CPU intensive tasks. There is a problem when you try to calculate the Fibonacci value for n-th number when `n >= 20` and `m >= 20`. So the code uses `memorization` strategy in Dynamic Programming with an algorythm to speed up the calculation from a minute long process (with GO! more with other languages, like Python or Node.js) to reduce it to miliseconds.

# The program

## Specify the n-th Fibonacci

You can change `line 10` to specify what is the `m` for column, and `n` for row values respectively.

e.g. 
``` Go
var column, row uint = 18, 18
```
## Run the program

Move to the folder where the `gridMath.go` is located on you command line or bash.
Type `Go run .` to run the program and get the response. I have also calculated the sequence from 1st to 10th, you will see the responses on your console.



