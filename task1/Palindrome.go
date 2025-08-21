package task1

import (
	"strconv"
	"strings"
)

// 判断是否是回文
func VerifyPalindrome(param int64) bool {
	str := strconv.FormatInt(param, 10)
	var newStr strings.Builder
	strArray := []rune(str)
	for i := len(strArray) - 1; i > -1; i-- {
		newStr.WriteRune(strArray[i])
	}
	if newStr.String() == str {
		return true
	}
	return false
}
