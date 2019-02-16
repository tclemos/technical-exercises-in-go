package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(MergeSort(StringArrayToInt(os.Args[1:])))
}

// StringArrayToInt Converts an array of strings containing
// numbers into an array of integers
func StringArrayToInt(a []string) []int {
	var integers []int

	for _, element := range a {
		var number, err = strconv.Atoi(element)
		if err == nil {
			integers = append(integers, number)
		}
	}

	return integers
}

// MergeSort performs a merge sort in an array
func MergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}
	mid := (len(a)) / 2
	return MergeSortedArrays(MergeSort(a[:mid]), MergeSort(a[mid:]))
}

// MergeSortedArrays merges 2 sorted arrays into a sorted one
func MergeSortedArrays(a1 []int, a2 []int) []int {

	if len(a1) <= 0 {
		return a2
	}

	if len(a2) <= 0 {
		return a1
	}

	var merged []int

	a1Index := 0
	a2Index := 0

	for a1Index < len(a1) || a2Index < len(a2) {

		if a1[a1Index] < a2[a2Index] {
			merged = append(merged, a1[a1Index])
			a1Index++
		} else {
			merged = append(merged, a2[a2Index])
			a2Index++
		}

		if a1Index == len(a1) || a2Index == len(a2) {
			break
		}
	}

	for i := a1Index; i < len(a1); i++ {
		merged = append(merged, a1[i])
	}

	for i := a2Index; i < len(a2); i++ {
		merged = append(merged, a2[i])
	}

	return merged
}
