package main

import (
	"fmt"
)

func main() {

	//有效的括号
	result3 := isValidParent("([]){}") // 测试括号匹配
	fmt.Println("括号匹配结果：", result3)

}

// 有效的括号
func isValidParent(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if open, exists := bracketMap[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
