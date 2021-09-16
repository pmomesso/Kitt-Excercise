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

func amountOfPeriodRoundedUp(minutes int, periodLength int) int {
	amountQuotient := float64(minutes) / float64(periodLength)
	amountCeil := int(math.Ceil(amountQuotient))

	return amountCeil
}

func amountOfWeeksRoundedUp(minutes int) int {
	return amountOfPeriodRoundedUp(minutes, WEEK_MINUTES)
}

func amountOfHoursRoundedUp(minutes int) int {
	return amountOfPeriodRoundedUp(minutes, HOUR_MINUTES)
}

func greaterThanDay(minutes int) bool {
	return minutes > DAY_MINUTES
}

func preciselyOneDay(minutes int) bool {
	return minutes == DAY_MINUTES
}

func greaterThanMinute(minutes int) bool {
	return minutes > 1
}

func GetPrice(minutes int) int {
	if greaterThanDay(minutes) {
		return amountOfWeeksRoundedUp(minutes) * WEEKLY_FARE
	}

	if preciselyOneDay(minutes) {
		return 1 * DAILY_FARE
	}

	if greaterThanMinute(minutes) {
		return amountOfHoursRoundedUp(minutes) * HOURLY_FARE
	}

	return minutes * 2
}
