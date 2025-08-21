package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//是不是回文数
	var param int64 = -121
	result2 := verifyPalindrome(param)
	if result2 {
		fmt.Println(param, "是回文数")
	} else {
		fmt.Println(param, "不是回文数")
	}

}

// 判断是否是回文
func verifyPalindrome(param int64) bool {
	str := strconv.FormatInt(param, 10)
	var newStr strings.Builder
	strArray := []rune(str)
	for i := len(strArray) - 1; i > -1; i-- {
		newStr.WriteRune(strArray[i])
	}
	if strings.EqualFold(newStr.String(), str) {
		return true
	}
	return false
}
