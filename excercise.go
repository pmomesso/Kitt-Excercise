package kittexcercise

const DAILY_FARE = 60
const WEEKLY_FARE = 105

const DAY_MINUTES = 24 * 60
const WEEK_MINUTES = 24 * 60 * 7

func isWeeks(minutes int) bool {
	return minutes%WEEK_MINUTES == 0
}

func amountOfWeeks(minutes int) int {
	amount := minutes / WEEK_MINUTES
	if amount == 0 {
		return 1
	}
	return amount
}

func GetPrice(minutes int) int {
	if minutes == 0 {
		return 0
	}

	if isWeeks(minutes) || minutes > DAY_MINUTES {
		return amountOfWeeks(minutes) * WEEKLY_FARE
	}

	return DAILY_FARE
}
