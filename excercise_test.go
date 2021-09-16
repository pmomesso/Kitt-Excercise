package kittexcercise

import (
	"fmt"
	"testing"
)

func TestGetPrice(t *testing.T) {
	var testCases = []struct {
		minutes, expected int
	}{
		{0, 0},
	}
	for _, testCase := range testCases {
		testCaseName := fmt.Sprintf("Minutes:%v,Price:%v", testCase.minutes, testCase.expected)
		t.Run(testCaseName, func(t *testing.T) {
			actualPrice := GetPrice(testCase.minutes)
			if actualPrice != testCase.expected {
				t.Errorf("Expected %v pounds, got %v pounds, for %v minutes", testCase.expected, actualPrice, testCase.minutes)
			}
		})
	}
}
