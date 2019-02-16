package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	arraytools "github.com/tclemos/technical-exercises-in-go/array-tools"
	twosum "github.com/tclemos/technical-exercises-in-go/arrays/two-sum/lib"
)

type exitCodeReturner func(code int)

var (
	returnExitCode exitCodeReturner = func(code int) {
		os.Exit(code)
	}
)

func main() {

	target, numbers, err := parseParameters()
	if err != nil {
		log.Print(err.Error())
		returnExitCode(1)
		return
	}

	i1, i2, v1, v2, err := twosum.Find(target, numbers)
	if err != nil {
		log.Print(err.Error())
		returnExitCode(1)
		return
	}

	log.Print(i1, i2, v1, v2)
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
		return target, nil, fmt.Errorf("Invalid target value: Target should be a number and \"%s\" is not a number", os.Args[1])
	}

	// Inject all the other arguments into the func toIntSlice to convert
	// all the string values into int values
	numbers = arraytools.ToIntSlice(os.Args[2:])
	if len(numbers) < 2 {
		return 0, nil, errors.New("The collection of numbers must have at least 2 numbers")
	}

	// Once everything is correct, returns the target and the numbers without error.
	return target, numbers, nil
}
