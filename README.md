# Dynamic Programming

## The goal

This repository is going to be a place where I would regualrly solve algorythmic math challenges, and update it. I've started this repo out of passion for writing a more useful practice code for `golang` and test my own limits. Later on I may add more languages in the mix. 

The coding challenges would require Dynamic Programming algorythms or strategies to solve. But for now I'm focusing on the following two strategies: 

### Memorization

I'm writing a function to handle the problems. It's successful but the normal solution is CPU intensive for large/big numbers. Therefore, it's essential to save common patterns of a virtual (imaginative) tree in `memory` so if we get to it again we return those previous calculations and save time.

e.g. create a mapping `map[SOME_TYPE]ANOTHER_TYPE` equevalant to an `Object {}` in JS types to memorize the patterns and refer to them later. 


## Tabulation

Tabulation is a similar process with **Memorization** except it is performed with a `Slice` a type of array without a cap limit in __GoLang__.

# Challenges so far

## The n-th Fibonacci Math Challenge

It can be found in the `/fibonacci-math` folder in this repo.
Refer to its folder specific `README.md` to learn more.

## The Grid Traveler Challenge

It can be found in the `/grid-traveler` folder in this repo.
Refer to its folder specific `README.md` to learn more.



