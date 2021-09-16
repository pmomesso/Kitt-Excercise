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
		{DAY_MINUTES, DAILY_TARIFF},
		{2 * WEEK_MINUTES, 2 * WEEKLY_TARIFF},
		{2 * DAY_MINUTES, WEEKLY_TARIFF},
		{2 * HOUR_MINUTES, 2 * HOURLY_TARIFF},
		{30, 30 * MINUTE_TARIFF},
		{1, MINUTE_TARIFF},
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
