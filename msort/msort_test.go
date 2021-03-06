package msort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

type testCase struct {
	input  []interface{}
	output []interface{}
}

func runTestcases(tests []testCase, cmp Compare, t *testing.T) {
	for _, test := range tests {
		result := Sort(test.input, cmp)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Failed. Expected: %v, Got: %v", test.output, test.input)
		}
	}
}

func dataset(size int) (data []interface{}) {
	data = make([]interface{}, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	return data
}

func randomize(data []interface{}) []interface{} {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	return data
}

func reverse(data []interface{}) []interface{} {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}

func TestIntegerSort(t *testing.T) {

	testCases := []testCase{
		{
			input:  []interface{}{5, 4, 3, 2, 1},
			output: []interface{}{1, 2, 3, 4, 5},
		},
		{
			input:  []interface{}{1, 2, 3, 4},
			output: []interface{}{1, 2, 3, 4},
		},
		{
			input:  []interface{}{2, 1},
			output: []interface{}{1, 2},
		},
		{
			input:  []interface{}{3, 2, 4, 1, 5, 5, 6, 1},
			output: []interface{}{1, 1, 2, 3, 4, 5, 5, 6},
		},
		{
			input:  []interface{}{1, 1, 1, 1, 1, 1, 1, 1},
			output: []interface{}{1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			input:  []interface{}{1},
			output: []interface{}{1},
		},
		{
			input:  []interface{}{},
			output: []interface{}{},
		},
		{ // Large data set -- randomized
			input:  randomize(dataset(1000 * 1000)), // 1M data
			output: dataset(1000 * 1000),
		},
		{ // Large data set -- worstcase
			input:  reverse(dataset(100 * 100)), // 100K data
			output: dataset(100 * 100),
		},
	}
	runTestcases(testCases,
		func(a, b interface{}) int {
			v1 := a.(int)
			v2 := b.(int)
			if v1 == v2 {
				return 0
			} else if v1 < v2 {
				return -1
			}
			return 1
		},
		t)
}

func TestStringSort(t *testing.T) {
	testCases := []testCase{
		{
			input:  []interface{}{"b", "a"},
			output: []interface{}{"a", "b"},
		},
		{
			input:  []interface{}{"d", "c", "b", "a"},
			output: []interface{}{"a", "b", "c", "d"},
		},
		{
			input:  []interface{}{"A", "a", "AA", "aa"},
			output: []interface{}{"A", "AA", "a", "aa"},
		},
		{
			input:  []interface{}{"da", "ac", "zb", "a"},
			output: []interface{}{"a", "ac", "da", "zb"},
		},
		{
			input:  []interface{}{"a", "a", "a", "a"},
			output: []interface{}{"a", "a", "a", "a"},
		},
	}
	runTestcases(testCases,
		func(a, b interface{}) int {
			v1 := a.(string)
			v2 := b.(string)
			if v1 == v2 {
				return 0
			} else if v1 < v2 {
				return -1
			}
			return 1
		},
		t)
}
