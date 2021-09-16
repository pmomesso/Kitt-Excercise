package kittexcercise

import "math"

const DAILY_FARE = 60
const WEEKLY_FARE = 105

const DAY_MINUTES = 24 * 60
const WEEK_MINUTES = 24 * 60 * 7

func isWeeks(minutes int) bool {
	return minutes%WEEK_MINUTES == 0
}

func amountOfWeeksCeil(minutes int) int {
	amountQuotient := float64(minutes) / WEEK_MINUTES
	amountCeil := int(math.Ceil(amountQuotient))

	return amountCeil
}

func greaterThanDay(minutes int) bool {
	return minutes > DAY_MINUTES
}

func GetPrice(minutes int) int {
	if minutes == 0 {
		return 0
	}

	if isWeeks(minutes) || greaterThanDay(minutes) {
		return amountOfWeeksCeil(minutes) * WEEKLY_FARE
	}

	return DAILY_FARE
}
