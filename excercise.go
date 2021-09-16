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

func amountOfPeriodCeil(minutes int, periodLength int) int {
	amountQuotient := float64(minutes) / float64(periodLength)
	amountCeil := int(math.Ceil(amountQuotient))

	return amountCeil
}

func amountOfWeeksCeil(minutes int) int {
	return amountOfPeriodCeil(minutes, WEEK_MINUTES)
}

func amountOfHoursCeil(minutes int) int {
	return amountOfPeriodCeil(minutes, HOUR_MINUTES)
}

func greaterThanDay(minutes int) bool {
	return minutes > DAY_MINUTES
}

func preciselyDay(minutes int) bool {
	return minutes == DAY_MINUTES
}

func GetPrice(minutes int) int {
	if greaterThanDay(minutes) {
		return amountOfWeeksCeil(minutes) * WEEKLY_FARE
	}

	if preciselyDay(minutes) {
		return 1 * DAILY_FARE
	}

	return amountOfHoursCeil(minutes) * HOURLY_FARE
}
