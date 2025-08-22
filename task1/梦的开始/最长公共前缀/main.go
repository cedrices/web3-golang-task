package main

import (
	"fmt"
)

func main() {

	//最长公共前缀
	arrays2 := []string{"flow", "flower", "flight"}
	result4 := maxPrefix(arrays2)
	fmt.Println("最长公共前缀是：", result4)
}

func maxPrefix(arrays []string) string {
	if len(arrays) == 0 {
		return ""
	}

	prefix := arrays[0]

	for i := 1; i < len(arrays); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(arrays[i]) || prefix[j] != arrays[i][j] {
				if j == 0 {
					return "" // 如果第一个字符就不匹配，直接返回空字符串
				}
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}
