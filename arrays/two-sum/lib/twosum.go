package twosum

import "errors"

// Find tries to find two numbers in the provided numbers
// collection that when added up together will match the provided target
//
// Returns the indexes and the values of the numbers found or an error
func Find(target int, numbers []int) (i1 int, i2 int, v1 int, v2 int, err error) {

	possiblePairs := make(map[int]int)

	for currentIndex := 0; currentIndex < len(numbers); currentIndex++ {

		currentNumber := numbers[currentIndex]
		pairNumber := target - currentNumber

		if pairIndex, found := possiblePairs[currentNumber]; found {
			return pairIndex, currentIndex, pairNumber, currentNumber, nil
		}

		possiblePairs[pairNumber] = currentIndex
	}

	return 0, 0, 0, 0, errors.New("The collection contains no pair that matches the target value when added up")
}
