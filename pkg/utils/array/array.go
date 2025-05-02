package array

func Contains(array []string, x string) bool {
	for _, i := range array {
		if i == x {
			return true
		}
	}
	return false
}
