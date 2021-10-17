package utils

// receives flag ("odd", "even", "all") and counter
// determines if current count is even or odd
// returns bool
func SwitchTo(flag string, counter int64) bool {
	switch flag {
	case "odd":
		return counter % 2 != 0
	case "even":
		return counter % 2 == 0
	default:
		return true
	}
}
