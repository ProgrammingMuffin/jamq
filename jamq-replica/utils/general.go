package utils

func InArray(element interface{}, array []interface{}) bool {
	for _, val := range array {
		if element == val {
			return true
		}
	}
	return false
}
