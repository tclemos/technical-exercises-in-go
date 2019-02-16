package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMainWithCorrectParameters(T *testing.T) {
	// Arrange
	expectedResult := "0 1 2 7\n"
	os.Args = []string{"", "9", "2", "7", "11", "15"}

	var buffer bytes.Buffer
	log.SetOutput(&buffer)

	// Act
	main()

	// Assert
	result := buffer.String()

	if strings.HasSuffix(result, expectedResult) {
		T.Log("Test executed successfully.")
	} else {
		T.Errorf("\nInvalid result\nExpected it to ends with: %s\nFound: %s", strings.Replace(expectedResult, "\n", "", -1), result)
	}
}

func TestMainInvalidTarget(T *testing.T) {
	// Arrange
	expectedResult := "Invalid target value: Target should be a number and \"A\" is not a number\n"
	os.Args = []string{"", "A", "2", "7", "11", "15"}

	var exitCode int

	returnExitCode = func(code int) { exitCode = code }

	var buffer bytes.Buffer
	log.SetOutput(&buffer)

	// Act
	main()

	// Assert
	if exitCode != 1 {
		T.Errorf("Wrong exit code returned\nExpected: %d\nFound: %d", 1, exitCode)
	}

	result := buffer.String()

	if strings.HasSuffix(result, expectedResult) {
		T.Log("Test executed successfully.")
	} else {
		T.Errorf("\nInvalid result\nExpected it to ends with: %s\nFound: %s", strings.Replace(expectedResult, "\n", "", -1), result)
	}
}

func TestMainWithoutRequiredParameters(T *testing.T) {
	// Arrange
	expectedResult := "Invalid number of arguments\n"

	testCases := []struct {
		Input []string
	}{
		{[]string{""}},
		{[]string{"", "2"}},
		{[]string{"", "2", "2"}},
	}

	for _, testCase := range testCases {
		os.Args = testCase.Input

		var exitCode int

		returnExitCode = func(code int) { exitCode = code }

		var buffer bytes.Buffer
		log.SetOutput(&buffer)

		// Act
		main()

		// Assert
		if exitCode != 1 {
			T.Errorf("Wrong exit code returned\nExpected: %d\nFound: %d", 1, exitCode)
		}

		result := buffer.String()

		if strings.HasSuffix(result, expectedResult) {
			T.Log("Test executed successfully.")
		} else {
			T.Errorf("\nInvalid result\nExpected it to ends with: %s\nFound: %s", strings.Replace(expectedResult, "\n", "", -1), result)
		}
	}
}

func TestMainWithInsufficientNumbersToCompare(T *testing.T) {
	// Arrange
	expectedResult := "The collection of numbers must have at least 2 numbers\n"

	testCases := []struct {
		Input []string
	}{
		{[]string{"", "4", "1", "A"}},
		{[]string{"", "4", "1", "A", "B"}},
		{[]string{"", "4", "1", "A", ""}},
		{[]string{"", "4", "1", "", ""}},
	}

	for _, testCase := range testCases {
		os.Args = testCase.Input

		var exitCode int

		returnExitCode = func(code int) { exitCode = code }

		var buffer bytes.Buffer
		log.SetOutput(&buffer)

		// Act
		main()

		// Assert
		if exitCode != 1 {
			T.Errorf("Wrong exit code returned\nExpected: %d\nFound: %d", 1, exitCode)
		}

		result := buffer.String()

		if strings.HasSuffix(result, expectedResult) {
			T.Log("Test executed successfully.")
		} else {
			T.Errorf("\nInvalid result\nExpected it to ends with: %s\nFound: %s", strings.Replace(expectedResult, "\n", "", -1), result)
		}
	}
}
