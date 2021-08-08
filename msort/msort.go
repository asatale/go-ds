package msort

//Compare function returns
// -1 if first quantity is less than second
//  0 if first quantity is same as second
//  1 if first quantity is greater than second
type Compare func(interface{}, interface{}) int

//Sort does sorting on the data using merge-sort algorithm
func Sort(data []interface{}, cmp Compare) []interface{} {
	if len(data) == 0 {
		return []interface{}{}
	}
	return msortIterative(data, cmp)
}

// msortRecursive is recursive implementation for merge-sort algorithm
func msortRecursive(data []interface{}, cmp Compare) []interface{} {
	var fn func([]interface{}, int, int) []interface{}

	fn = func(data []interface{}, lo, hi int) []interface{} {
		if lo < hi {
			mid := int(lo + ((hi - lo) / 2))
			s1 := fn(data, lo, mid)
			s2 := fn(data, mid+1, hi)
			result := make([]interface{}, 0, (hi - lo))
			return merge(s1, s2, cmp, result)
		}
		return []interface{}{data[lo]}
	}
	return fn(data, 0, len(data)-1)
}

func msortIterative(data []interface{}, cmp Compare) []interface{} {
	result := make([]interface{}, 0, len(data))

	for interval := 1; interval < len(data); interval = interval * 2 {
		for i := 0; i+interval < len(data); i = i + 2*interval {
			lo := i
			mid := lo + interval
			hi := mid + interval
			if hi > len(data) {
				hi = len(data)
			}
			result = result[:0]
			result = merge(data[lo:mid], data[mid:hi], cmp, result)
			copy(data[lo:hi], result)
		}
	}
	return data
}

// merge performs merge for two sorted slices to produce sorted slice
func merge(s1, s2 []interface{}, cmp Compare, result []interface{}) []interface{} {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if cmp(s1[i], s2[j]) == 1 {
			result = append(result, s2[j])
			j = j + 1
		} else {
			result = append(result, s1[i])
			i = i + 1
		}
	}

	for i < len(s1) {
		result = append(result, s1[i])
		i = i + 1
	}

	for j < len(s2) {
		result = append(result, s2[j])
		j = j + 1
	}

	return result
}
