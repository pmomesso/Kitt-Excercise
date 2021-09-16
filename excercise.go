package kittexcercise

import "math"

const HOURLY_FARE = 22
const DAILY_FARE = 60
const WEEKLY_FARE = 105

const HOUR_MINUTES = 60
const DAY_MINUTES = 24 * 60
const WEEK_MINUTES = 24 * 60 * 7

func isExactlyPeriod(minutes int, periodLength int) bool {
	return minutes%periodLength == 0
}

func isWeeks(minutes int) bool {
	return isExactlyPeriod(minutes, WEEK_MINUTES)
}

func isDays(minutes int) bool {
	return isExactlyPeriod(minutes, DAY_MINUTES)
}

func amountOfWeeksCeil(minutes int) int {
	amountQuotient := float64(minutes) / WEEK_MINUTES
	amountCeil := int(math.Ceil(amountQuotient))

	return amountCeil
}

func greaterThanDay(minutes int) bool {
	return minutes > DAY_MINUTES
}

func amountOfDaysCeil(minutes int) int {
	amountQuotient := float64(minutes) / HOUR_MINUTES
	amount := int(math.Ceil(amountQuotient))

	return amount
}

func GetPrice(minutes int) int {
	if minutes == 0 {
		return 0
	}

	if isWeeks(minutes) || greaterThanDay(minutes) {
		return amountOfWeeksCeil(minutes) * WEEKLY_FARE
	}

	if isDays(minutes) {
		return DAILY_FARE
	}

	return amountOfDaysCeil(minutes) * HOURLY_FARE
}
