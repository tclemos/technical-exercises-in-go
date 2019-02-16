package arraytools

import "strconv"

// ToIntSlice converts a string slice into an int slice.
//
// All the non int compliant string values will be skipped
// and it will always return an instance of slice and never returns nil
func ToIntSlice(entryValues []string) []int {

	if entryValues == nil {
		return make([]int, 0)
	}

	result := make([]int, 0)

	for _, entryValue := range entryValues {
		if convertedValue, err := strconv.Atoi(entryValue); err == nil {
			result = append(result, convertedValue)
		}
	}

	return result
}
