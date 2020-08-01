package main

// arr: [ -10 -5 -3 0 2 3 6 8 10 ]
// curIdx: 6
// arr[curIdx]: 6

// arr: []
// curIdx: 0

// arr: [ -2 -1  ]

// func magicIndex(arr []int) int {
// 	for curIdx := 0; curIdx < len(arr); curIdx++ {
// 		if arr[curIdx] > curIdx {
// 			curIdx = arr[curIdx]
// 			continue
// 		} else if arr[curIdx] == curIdx {
// 			return curIdx
// 		}
// 	}
// 	return -1
// }

// arr: [ -10 -5 -3 0 2 3 6 8 10 ]
// low: 2
// high: 6
// arr[low]: -3
// arr[high]: 6

// func magicIndex(arr []int) int {
// 	low := 0
// 	high := len(arr) - 1
// 	for low < high {
// 		if arr[low] == low {
// 			return low
// 		} else if arr[low] > low {
// 			low = arr[low]
// 		} else {
// 			low++
// 		}
// 		if arr[high] == high {
// 			return high
// 		} else if arr[high] < high {
// 			high = arr[high]
// 		} else {
// 			high--
// 		}
// 	}
// 	return -1
// }

func magicIndex(arr []int) int {
	return magicIndexHelper(arr, 0, len(arr)-1)
}

func magicIndexHelper(arr []int, start int, end int) int {
	if end-start < 0 {
		return -1
	}
	mid := (end-start)/2 + start
	if arr[mid] == mid {
		return mid
	} else if arr[mid] < mid {
		return magicIndexHelper(arr, start+1, end)
	}
	return magicIndexHelper(arr, start, mid-1)
}
