package mergesort

// AscendingSort sorts a slice of integer in ascending order
func AscendingSort(unorderedItems []int) []int {

	if unorderedItems == nil {
		return []int{}
	}

	// If the array has a size of one, it is already sorted.
	if len(unorderedItems) <= 1 {
		return unorderedItems
	}

	// find the half of the slice
	mid := (len(unorderedItems)) / 2

	// "{startIndex}:{endIndex}" is the symbol used to retrieve a set of elements in a sequence from a slice
	//
	// The startIndex should represents the index where the set of items starts and this is inclusive, this
	// means that the index provided in startIndex will be retrieved, when omitted there 0 index will be
	// considered
	//
	// The endIndex represents the index where the set of items ends and this is NOT inclusive, this means
	// that the index provided in endIndex will NOT be retrieved, when omitted the last index(len(arr)-1) will
	// be considered
	//
	// Example: arr := [1, 2, 3, 4, 5]
	// arr[0:0] = []
	// arr[1:3] = [2, 3]
	// arr[2:5] = [3, 4, 5]
	// arr[0:4] = [1, 2, 3, 4]
	// arr[0:5] = [1, 2, 3, 4, 5]

	// Gets the first half of elements from the slice of integers from index 0 until the middle index
	leftUnorderedItems := unorderedItems[:mid]
	// Gets the second half of elements from the slice of integers from the middle middle until the last index
	rightUnorderedItems := unorderedItems[mid:]

	// Call recursively the AscendingSort function to order the elements
	leftSortedItems := AscendingSort(leftUnorderedItems)
	rightSortedItems := AscendingSort(rightUnorderedItems)

	// Merge the ordered elements into one slice of integers
	sortedItems := AscendingMerge(leftSortedItems, rightSortedItems)

	return sortedItems
}

// AscendingMerge merges 2 sorted slices of integers into
// a sorted in ascending order
//
// leftOrderedItems and rightOrderedItems must be already ordered in ascendent order
func AscendingMerge(leftOrderedItems []int, rightOrderedItems []int) []int {

	if len(rightOrderedItems) == 0 {
		// Otherwise, when the right slice has no elements we return the left slice
		return leftOrderedItems
	}

	// Prepare a new slice to contain the ordered items from both right and left slices
	var merged []int

	// keep track of the slice positions
	leftOrderedItemsIndex := 0
	rightOrderedItemsIndex := 0

	for leftOrderedItemsIndex < len(leftOrderedItems) || rightOrderedItemsIndex < len(rightOrderedItems) {
		// While the indexes does not reach the end of the slice, we keep navigating in both slices

		if leftOrderedItems[leftOrderedItemsIndex] < rightOrderedItems[rightOrderedItemsIndex] {
			// If the current left item is smaller than the current right item
			// we add the left item to the merge slice
			merged = append(merged, leftOrderedItems[leftOrderedItemsIndex])
			// and move the left index to the next position
			leftOrderedItemsIndex++
		} else {
			// If the current right item is greater than or equal to the current left item
			// we add the right item to the merge slice
			merged = append(merged, rightOrderedItems[rightOrderedItemsIndex])
			// and move the right index to the next position
			rightOrderedItemsIndex++
		}

		if leftOrderedItemsIndex == len(leftOrderedItems) || rightOrderedItemsIndex == len(rightOrderedItems) {
			// If the left index is greater than the left slice or the right index is
			// greater than the right slice we stop navigating in the indexes, because now
			// we just need to get the numbers from the slice that still has numbers
			break
		}
	}

	for i := leftOrderedItemsIndex; i < len(leftOrderedItems); i++ {
		// while the left index is small than the left slice size
		// add them to the merge slice
		merged = append(merged, leftOrderedItems[i])
	}

	for i := rightOrderedItemsIndex; i < len(rightOrderedItems); i++ {
		// while the right index is small than the right slice size
		// add them to the merge slice
		merged = append(merged, rightOrderedItems[i])
	}

	// return the merged slice with all the numbers provided by left
	// and right slices, ordered in an ascending order
	return merged
}
