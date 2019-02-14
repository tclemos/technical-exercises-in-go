package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	twosum "github.com/tclemos/technical-exercises-in-go/arrays/two-sum/lib"
)

func main() {

	target, numbers, err := parseParameters()
	if err != nil {
		log.Fatal(err.Error())
	}

	i1, i2, v1, v2, err := twosum.Find(target, numbers)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(i1, i2, v1, v2)
}

// ParseParameters read all the command line arguments
// and tries to parse all the values into target and numbers
//
// The first argument provided is going to be considered
// as the target and the rest of the arguments are going
// be the collection of numbers.
//
// If something is not compliant with the requirements
// an error will be thrown with the reason
//
// It returns the target and numbers or an error
func parseParameters() (target int, numbers []int, err error) {

	// Checks the number of minimum arguments required
	//
	// Since the application requires the user to provide at
	// least 3 arguments, that will represent the target and
	// a collection with at least 2 numbers.
	//
	// Here we are just asking if the os.Args length is greater
	// than 4 because the first position of Args will always
	// provide the value that represents the path where the
	// application is running, so the arguments provided will
	// be available from the position number 1 in the os.Args
	// array. e.g os.Args[1:]
	if len(os.Args) < 4 {
		return 0, nil, errors.New("Invalid number of arguments")
	}

	// Tries to convert a string into an int that represents the Target
	//
	// strconv.Atoi tries to convert a string into an int in base 10
	target, err = strconv.Atoi(os.Args[1])
	if err != nil {
		return target, nil, fmt.Errorf("Target not valid: %s", err.Error())
	}

	// Inject all the other arguments into the func toIntSlice to convert
	// all the string values into int values
	numbers = toIntSlice(os.Args[2:])
	if len(numbers) < 1 {
		return 0, nil, errors.New("The collection of numbers should have at least 2 numbers")
	}

	// Once everything is correct, returns the target and the numbers without error.
	return target, numbers, nil
}

// ToIntSlice converts a string slice into an int slice.
//
// All the non int compliant string values will be skipped
// and it will always return an instance of slice and never returns nil
func toIntSlice(entryValues []string) []int {

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
