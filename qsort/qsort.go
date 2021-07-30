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
  qsort(data, 0, len(data)-1, cmp)
}

func qsort(data []interface{}, lo, hi int, cmp Compare) {
  if lo < hi {
    p := partition(data, lo, hi, cmp)
    qsort(data, lo, p-1, cmp)
    qsort(data, p+1, hi, cmp)
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
