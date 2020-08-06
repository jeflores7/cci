package main

// Run modified binary search to find the true 0 index
// On first iteration, compare mid to current 0 index to determine whether the bounds must go left or right
// Then treat it as normal bin search
// Run modified binary search with knowledge of true 0 index

// arr: [ 15, 16, 19, 20, 25, 30, 32, 37, 39, 42, 45, 55, 1, 3, 4, 5, 7, 10, 14 ], value: 1
// trueZero: 12

func searchRotated(arr []int, value int) int {
	trueZero := getTrueZero(arr, 0, len(arr)-1)
	return binSearch(arr, value, trueZero, trueZero-1)
}

// arr: [ 15, 16, 19, 20, 25, 30, 32, 37, 39, 42, 45, 55, 1, 3, 4, 5, 7, 10, 14 ]

// start: 0, end: 18
// mid: 9
// arr[mid]: 42, arr[mid+1]: 45, arr[mid-1]: 39
// arr[start]: 15

// start: 10, end: 18
// mid: 14
// arr[mid]: 4, arr[mid+1]: 5, arr[mid-1]: 3
// arr[start]: 45

// start: 10, end: 13
// mid: 11
// arr[mid]: 55, arr[mid+1]: 1, arr[mid-1]: 45
// arr[start]: 45

func getTrueZero(arr []int, start int, end int) int {
	mid := (start + end) / 2
	if arr[mid] < arr[mid-1] {
		return mid
	} else if arr[mid] > arr[mid+1] {
		return mid + 1
	}
	if arr[mid] > arr[start] {
		return getTrueZero(arr, mid+1, end)
	}
	return getTrueZero(arr, start, mid-1)
}

// arr: [ 15, 16, 19, 20, 25, 30, 32, 37, 39, 42, 45, 55, 1, 3, 4, 5, 7, 10, 14 ]

// value: 1, start: 12, end: 11
// mid: 3
// arr[mid]: 20

// value: 1, start: 12, end: 2
// mid: 16
// arr[mid]: 7

// value: 1, start: 12, end: 15
// mid:

func binSearch(arr []int, value int, start int, end int) int {
	mid := getIndex(len(arr), start, end)
	if arr[mid] == value {
		return mid
	}
	if arr[mid] > value {
		if mid == 0 {
			return binSearch(arr, value, start, len(arr)-1)
		}
		return binSearch(arr, value, start, mid-1)
	}
	if mid == len(arr)-1 {
		return binSearch(arr, value, 0, end)
	}
	return binSearch(arr, value, mid+1, end)
}

// arrLen: 19, start: 12, end: 11
// half: 9
// return 21 % 18 -> 3

// arrLen: 19, start: 12, end: 2
// half: 4
// return 16 % 18 -> 16

// arrLen: 19, start: 12, end: 15
// half: 13
// return 13

func getIndex(arrLen int, start int, end int) int {
	if start > end {
		half := ((arrLen - start) + end) / 2
		return (start + half) % arrLen
	}
	return (start + end) / 2
}
