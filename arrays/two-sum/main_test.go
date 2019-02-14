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
