package kittexcercise

const HOURLY_FARE = 60

func GetPrice(minutes int) int {
	if minutes == 0 {
		return 0
	}

	return HOURLY_FARE
}
