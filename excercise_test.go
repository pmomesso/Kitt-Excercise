package kittexcercise

import (
	"errors"
	"fmt"
	"testing"
)

func TestGetPrice(t *testing.T) {
	var testCases = []struct {
		minutes, expected int
		err               error
	}{
		{0, 0, nil},
		{DAY_MINUTES, DAILY_TARIFF, nil},
		{2 * WEEK_MINUTES, 2 * WEEKLY_TARIFF, nil},
		{2 * DAY_MINUTES, WEEKLY_TARIFF, nil},
		{2 * HOUR_MINUTES, 2 * HOURLY_TARIFF, nil},
		{30, 30 * MINUTE_TARIFF, nil},
		{1, MINUTE_TARIFF, nil},
		{-1, 0, errors.New("")},
	}
	for _, testCase := range testCases {
		testCaseName := fmt.Sprintf("Minutes:%v,Price:%v", testCase.minutes, testCase.expected)
		t.Run(testCaseName, func(t *testing.T) {
			if testCase.err == nil {
				actualPrice, err := GetPrice(testCase.minutes)
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				} else if actualPrice != testCase.expected {
					t.Errorf("Expected %v pounds, got %v pounds, for %v minutes", testCase.expected, actualPrice, testCase.minutes)
				}
			} else {
				_, err := GetPrice(testCase.minutes)
				if err == nil {
					t.Errorf("Expected error to be %v, got nil error", testCase.err)
				}
			}
		})
	}
}
