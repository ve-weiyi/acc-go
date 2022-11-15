package utils

func If(isTrue bool, a, b interface{}) interface{} {
	if isTrue {
		return a
	}
	return b
}

// CheckData  []interface{}和 ...interface{} 有很大区别
func CheckData(data []interface{}) interface{} {
	if len(data) > 0 {
		return data[0]
	}
	return "No data returned"
}
