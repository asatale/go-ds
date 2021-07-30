package qsort

//Compare function returns
// -1 if first quantity is less than second
//  0 if first quantity is same as second
//  1 if first quantity is greater than second
type Compare func(interface{}, interface{}) int

//Sort does inplace sorting on the data using quick-sort algorithm
func Sort(data []interface{}, cmp Compare) {
	if len(data) < 2 {
		return
	}
	//qsortRecursive(data, 0, len(data)-1, cmp)
	qsortIterative(data, 0, len(data)-1, cmp)
}

// qsortRecursive is recursive implementation for quick-sort algorithm
func qsortRecursive(data []interface{}, lo, hi int, cmp Compare) {
	if lo < hi {
		p := partition(data, lo, hi, cmp)
		qsortRecursive(data, lo, p-1, cmp)
		qsortRecursive(data, p+1, hi, cmp)
	}
}

// qsortIterative is an iterative implementation for quick-sort algorithm
func qsortIterative(data []interface{}, lo, hi int, cmp Compare) {
	type qIteration struct {
		lo, hi int
	}
	iter := make([]qIteration, 0)
	iter = append(iter, qIteration{lo: lo, hi: hi})
	for len(iter) > 0 {
		ii := iter[len(iter)-1]
		iter = iter[:len(iter)-1]
		if ii.lo < ii.hi {
			p := partition(data, ii.lo, ii.hi, cmp)
			iter = append(iter, qIteration{lo: p + 1, hi: ii.hi}, qIteration{lo: ii.lo, hi: p - 1})
		}
	}
}

// partion algorithm as described in https://en.wikipedia.org/wiki/Quicksort
func partition(data []interface{}, lo, hi int, cmp Compare) int {
	pivot := hi
	loc := lo
	for i := lo; i <= hi; i++ {
		if cmp(data[i], data[pivot]) == -1 {
			data[loc], data[i] = data[i], data[loc]
			loc = loc + 1
		}
	}
	data[loc], data[pivot] = data[pivot], data[loc]
	return loc
}
