package kittexcercise

const HOURLY_FARE = 60

func GetPrice(minutes int) int {
	if minutes == 0 {
		return 0
	}

	if minutes%(24*60*7) == 0 {
		return minutes / (24 * 60 * 7) * 105
	}

	return HOURLY_FARE
}
