package utils

//
// StringContained
// @Description: 判断字符串是否在字符串切片中
// @param strArr
// @param str
// @return int
//
func StringContained(strArr []string, str string) int {
	for idx, s := range strArr {
		if s == str {
			return idx
		}
	}
	return -1
}
