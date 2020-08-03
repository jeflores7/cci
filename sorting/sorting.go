package sorting

func swap(arr []int, i int, j int) {
	if i != j {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	}
}

func getMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

func bubbleSort(arr []int) {
	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				swap(arr, i-1, i)
				swapped = true
			}
		}
	}
}

func selectionSort(arr []int) {
	for tracker := 0; tracker < len(arr); tracker++ {
		min := arr[tracker]
		minIdx := tracker
		for i := tracker; i < len(arr); i++ {
			if arr[i] < min {
				min = arr[i]
				minIdx = i
			}
		}
		swap(arr, tracker, minIdx)
	}
}

func mergeSort(arr []int) {
	mergeSortHelper(arr, make([]int, len(arr)), 0, len(arr)-1)
}

func mergeSortHelper(arr []int, helper []int, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSortHelper(arr, helper, start, mid)
	mergeSortHelper(arr, helper, mid+1, end)
	merge(arr, helper, start, mid, end)
}

func merge(arr []int, helper []int, start int, mid int, end int) {
	for i := start; i <= end; i++ {
		helper[i] = arr[i]
	}

	left := start
	right := mid + 1
	for i := start; i <= end; i++ {
		if left > mid {
			// remaining right are already in correct place
			break
		} else if right > end {
			arr[i] = helper[left]
			left++
		} else if helper[left] <= helper[right] {
			arr[i] = helper[left]
			left++
		} else {
			arr[i] = helper[right]
			right++
		}
	}
}

func quickSort(arr []int) {
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, start int, end int) {
	if start >= end {
		return
	}
	pivotPos := partition(arr, start, end)
	quickSortHelper(arr, start, pivotPos-1)
	quickSortHelper(arr, pivotPos, end)
}

func partition(arr []int, start int, end int) int {
	pivot := arr[(start+end)/2]
	i := start
	j := end
	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			swap(arr, i, j)
			i++
			j--
		}
	}
	return i
}

func radixSortBucketsLSD(arr []int) {
	buckets := make([][]int, 10)
	sigDigit := 1
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		bucket := (arr[i] / sigDigit) % 10
		buckets[bucket] = append(buckets[bucket], arr[i])
		if arr[i] > max {
			max = arr[i]
		}
	}
	radixSortLSD(arr, buckets, sigDigit*10, max)
}

func radixSortLSD(arr []int, buckets [][]int, sigDigit int, max int) {
	for {
		if max/sigDigit == 0 {
			break
		}
		nextB := make([][]int, 10)
		for digit := 0; digit < len(buckets); digit++ {
			curBucket := buckets[digit]
			for i := 0; i < len(curBucket); i++ {
				bucket := (curBucket[i] / sigDigit) % 10
				nextB[bucket] = append(nextB[bucket], curBucket[i])
			}
		}
		buckets = nextB
		sigDigit = sigDigit * 10
	}
	arrIdx := 0
	for b := 0; b < len(buckets); b++ {
		for i := 0; i < len(buckets[b]); i++ {
			arr[arrIdx] = buckets[b][i]
			arrIdx++
		}
	}
}

func radixSortCountingLSD(arr []int) {
	max := getMax(arr)
	for sigDigit := 1; max/sigDigit > 0; sigDigit *= 10 {
		radixSortCounting(arr, sigDigit)
	}
}

func radixSortCounting(arr []int, sigDigit int) {
	lenArr := len(arr)
	out := make([]int, lenArr)
	lenCount := 10
	count := make([]int, lenCount)
	for i := 0; i < len(arr); i++ {
		count[arr[i]/sigDigit%10]++
	}

	for i := 1; i < lenCount; i++ {
		count[i] += count[i-1]
	}

	for i := lenArr - 1; i >= 0; i-- {
		out[count[arr[i]/sigDigit%10]-1] = arr[i]
		count[arr[i]/sigDigit%10]--
	}

	for i := 0; i < lenArr; i++ {
		arr[i] = out[i]
	}
}

func radixSortBucketsMSD(arr []int) {
	max := getMax(arr)
	sigDigit := 1
	for max/(sigDigit*10) > 0 {
		sigDigit = sigDigit * 10
	}
	sorted := radixSortMSD(arr, sigDigit)
	for i := range sorted {
		arr[i] = sorted[i]
	}
}

func radixSortMSD(arr []int, sigDigit int) []int {
	if len(arr) == 1 || sigDigit == 0 {
		return arr
	}
	buckets := make([][]int, 10)
	sorted := make([]int, 0, len(arr))
	for _, value := range arr {
		bIdx := (value / sigDigit) % 10
		buckets[bIdx] = append(buckets[bIdx], value)
	}
	sigDigit = sigDigit / 10
	for i := range buckets {
		if len(buckets[i]) == 0 {
			continue
		}
		sorted = append(sorted, radixSortMSD(buckets[i], sigDigit)...)
	}
	return sorted
}

func countingSort(arr []int) {
	max := getMax(arr)
	count := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	idx := 0
	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			arr[idx] = i
			idx++
		}
	}
}
