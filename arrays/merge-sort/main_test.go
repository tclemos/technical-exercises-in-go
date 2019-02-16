package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMain(T *testing.T) {
	// Arrange
	expectedResult := "[-3 -2 -1 0 1 2 3]\n"
	os.Args = []string{"", "3", "1", "-1", "-3", "0", "2", "-2"}

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
