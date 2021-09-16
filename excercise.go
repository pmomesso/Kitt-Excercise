package kittexcercise

import (
	"errors"
	"math"
)

const MINUTE_TARIFF = 2
const HOURLY_TARIFF = 22
const DAILY_TARIFF = 60
const WEEKLY_TARIFF = 105

const HOUR_MINUTES = 60
const DAY_MINUTES = 24 * 60
const WEEK_MINUTES = 24 * 60 * 7

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

func greaterThanOneHour(minutes int) bool {
	return minutes > HOUR_MINUTES
}

func preciselyOneDay(minutes int) bool {
	return minutes == DAY_MINUTES
}

func invalidMinutesAmount(minutes int) bool {
	return minutes < 0
}

func GetPrice(minutes int) (int, error) {
	if invalidMinutesAmount(minutes) {
		return -1, errors.New("minutes amount was negative")
	}

	if greaterThanDay(minutes) {
		return amountOfWeeksRoundedUp(minutes) * WEEKLY_TARIFF, nil
	}

	if preciselyOneDay(minutes) {
		return 1 * DAILY_TARIFF, nil
	}

	if greaterThanOneHour(minutes) {
		return amountOfHoursRoundedUp(minutes) * HOURLY_TARIFF, nil
	}

	return minutes * MINUTE_TARIFF, nil
}
