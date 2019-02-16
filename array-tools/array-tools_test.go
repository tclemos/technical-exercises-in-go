package arraytools

import "testing"

func TestToIntSlice(T *testing.T) {

	testCases := []struct {
		Numbers  []string
		Expected []int
	}{
		{[]string{"1", "2", "3"}, []int{1, 2, 3}},
		{[]string{"1", "", "a", "2", "b", "3"}, []int{1, 2, 3}},
		{nil, []int{}},
		{[]string{}, []int{}},
		{[]string{""}, []int{}},
		{[]string{"a"}, []int{}},
		{[]string{"1"}, []int{1}},
	}

NextTestCase:
	for _, testCase := range testCases {
		result := ToIntSlice(testCase.Numbers)

		if len(result) != len(testCase.Expected) {
			T.Errorf("Wrong result size\nExpected: %d\nFound: %d", len(testCase.Expected), len(result))
			continue
		}

		for i, n := range result {
			if n != testCase.Expected[i] {
				T.Errorf("Wrong result\nExpected: %v\nFound: %v", testCase.Expected, result)
				continue NextTestCase
			}
		}
	}
}
