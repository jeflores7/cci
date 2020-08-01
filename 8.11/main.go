package main

import "strconv"

// Infinite number of:
// 25
// 10
// 5
// 1

// Calculate # of ways to represent n cents

// Take best case
// Take largest amount of largest coin
// Repeat for next coin on n - (previous coins taken)
// Repeat for next coin…
// Remainder are all pennies
// From largest coin, decrement taken by 1 and redo the problem until largest coin is 0
// Take largest amount of next coin on n - (previous coins taken)
// Repeat for next coin …
// Remainder are all pennies
// From next largest coin, decrememnt ....
// …

// n: 12

// func countWays(n int) int {
// 	return countWaysRecursive(n, []int{25, 10, 5, 1}, make(map[int]int))
// }

// n: 12, coins: [ 25, 10, 5, 1 ], memo: { 12: 4, 10: 3, 5: 1  } // return 4
// found: false, bestCase: { 25: 0, 10: 1, 5: 0, 1: 2  }, remainder: 0, coin: _
// numWays: 1 + 3
// i: 2, coin: 5, totalOfThisCoin: 0, remainingCoins: [ 5, 1 ]

// n: 10, coins: [ 5, 1 ], memo: { 10: 3, 5: 1 } // return 3
// found: false, bestCase: { 5: 2, 1: 0 }, remainder: 0, coin: _
// numWays: 1 + 1 + 1
// i: 0, coin: 1, totalOfThisCoin: 0, remainingCoins: [ 1 ]

// n: 5, coins: [ 1 ], memo: { 10: 1, 5: 1 } // return 1
// found: false, bestCase: { 1: 5 }, remainder: 0, coin: _
// numWays: 1
// i: 0, coin: 1

// n: 10, coins: [ 1 ], memo: { 10: 1 } // return 1
// found: false, bestCase: { 1: 10 }, remainder: 0, coin: _
// numWays: 1
// i: 1, coin: _

// func countWaysRecursive(n int, coins []int, memo map[int]int) int {
// 	if _, found := memo[n]; found {
// 		return memo[n]
// 	}
// 	bestCase := make(map[int]int)
// 	remainder := n
// 	for _, coin := range coins {
// 		bestCase[coin] = remainder / coin
// 		remainder = remainder - (bestCase[coin] * coin)
// 	}
// 	numWays := 1
// 	for i, coin := range coins {
// 		if coin == 1 {
// 			continue
// 		}
// 		totalOfThisCoin := bestCase[coin]
// 		remainingCoins := coins[i+1:]
// 		for totalOfThisCoin > 0 {
// 			numWays += countWaysRecursive(totalOfThisCoin*coin, remainingCoins, memo)
// 			totalOfThisCoin--
// 		}
// 	}
// 	memo[n] = numWays
// 	return memo[n]
// }

func countWays(n int) int {
	return countWaysRecursive([]int{25, 10, 5, 1}, n, 0, make(map[string]int))
}

func countWaysRecursive(coins []int, n int, idx int, memo map[string]int) int {
	if idx == len(coins)-1 {
		// 1 way to represent any value using only pennies
		return 1
	}
	mapKey := strconv.Itoa(n) + strconv.Itoa(idx)
	if _, found := memo[mapKey]; found {
		return memo[mapKey]
	}
	ways := 0
	for totCoin := 0; totCoin*coins[idx] < n; totCoin++ {
		ways += countWaysRecursive(coins, n-(totCoin*coins[idx]), idx+1, memo)
	}
	memo[mapKey] = ways
	return ways
}
