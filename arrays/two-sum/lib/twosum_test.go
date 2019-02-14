package twosum

import (
	"errors"
	"testing"
)

func TestFindTwoSumSuccessCases(t *testing.T) {

	// Arrange
	testCases := []struct {
		Target        int
		Numbers       []int
		ExpectedI1    int
		ExpectedI2    int
		ExpectedV1    int
		ExpectedV2    int
		ExpectedError error
	}{
		{9, []int{2, 7, 11, 15}, 0, 1, 2, 7, nil},
		{9, []int{2, 6, 11, 15}, 0, 0, 0, 0, errors.New("The collection contains no pair that matches the target value when added up")},
		{8, []int{9, 4, 3, 7, 5, 2, 0}, 2, 4, 3, 5, nil},
	}

	for _, testCase := range testCases {

		i1, i2, v1, v2, err := Find(testCase.Target, testCase.Numbers)

		if testCase.ExpectedError != nil {
			expectedErrorMessage := testCase.ExpectedError.Error()

			if err == nil {
				t.Errorf("Missing error message.\nExpected: %s", expectedErrorMessage)
			} else if errorMessage := err.Error(); expectedErrorMessage != errorMessage {
				t.Errorf("Wrong error message.\nExpected: %s\nReceived: %s", expectedErrorMessage, errorMessage)
			} else {
				t.Log("Wrong message provided successfully")
			}
		}

		if i1 != testCase.ExpectedI1 || i2 != testCase.ExpectedI2 || v1 != testCase.ExpectedV1 || v2 != testCase.ExpectedV2 {
			t.Errorf("\nInvalid result\nExpected: i1 = %d, i2 = %d, v1 = %d, i2 = %d \nReceived: i1 = %d, i2 = %d, v1 = %d, i2 = %d",
				testCase.ExpectedI1, testCase.ExpectedI2, testCase.ExpectedV1, testCase.ExpectedV2, i1, i2, v1, v2)
		} else {
			t.Log("Indexes and values provided successfully")
		}
	}
}
