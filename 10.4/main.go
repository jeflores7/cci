package main

// Start by searching using powers of 2 as our indexes
// Once we find an element larger than the target value or receive a -1, we do a binary search on the range exp/2 + 1 to exp-1, i.e, the last iterations index plus to to the current iterations index minus one

type list struct {
	data []int
}

func (l list) elementAt(i int) int {
	if i >= len(l.data) {
		return -1
	}
	return l.data[i]
}

func searchListy(l list, value int) int {
	if value < 0 {
		return -1
	}
	start := 0
	end := 2
	el := l.elementAt(end)
	for el != -1 && el < value {
		start = end + 1
		end = end * 2
		el = l.elementAt(end)
	}
	if el == -1 && end < value {
		return -1
	}
	return binSearch(l, value, start, end)
}
func binSearch(l list, value int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	el := l.elementAt(mid)
	if el == value {
		return mid
	}
	if el == -1 || el > value {
		return binSearch(l, value, start, mid-1)
	}
	return binSearch(l, value, mid+1, end)
}
