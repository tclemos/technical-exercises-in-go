package mergesort

import "testing"

func TestAscendingSort(T *testing.T) {

	testCases := []struct {
		Input    []int
		Expected []int
	}{
		{nil, []int{}},       //nil
		{[]int{}, []int{}},   // empty
		{[]int{1}, []int{1}}, // one element
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},             // already ordered
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},             // reverse ordered
		{[]int{2, 1, 2}, []int{1, 2, 2}},                               // dupped item
		{[]int{1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1}},       // repeated item
		{[]int{3, 2, 1, 0, -1, -2, -3}, []int{-3, -2, -1, 0, 1, 2, 3}}, // negative numbers
		{[]int{1, 2, 3, 4, 1, 2, 3, 4}, []int{1, 1, 2, 2, 3, 3, 4, 4}}, // equal half
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, []int{1, 1, 2, 2, 3, 3, 4, 4}}, // mirrored
	}

NextTestCase:
	for _, testCase := range testCases {
		result := AscendingSort(testCase.Input)

		if len(result) != len(testCase.Expected) {
			T.Errorf("Wrong result\nExpected: %v\nFound: %v", testCase.Expected, result)
		}

		for i, n := range result {
			if n != testCase.Expected[i] {
				T.Errorf("Wrong result\nExpected: %v\nFound: %v", testCase.Expected, result)
				continue NextTestCase
			}
		}
	}
}
