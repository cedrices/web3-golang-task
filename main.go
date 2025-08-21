package main

import (
	"fmt"

	"github.com/cedrices/web3-golang-task/task1"
)

func main() {
	//只出现一次的数字
	arrays := [11]int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5}
	result1 := task1.ArrayFindSingle(arrays)
	fmt.Println("只出现一次的数字是：", result1)

	//是不是回文数
	var param int64 = -121
	result2 := task1.VerifyPalindrome(param)
	if result2 {
		fmt.Println(param, "是回文数")
	} else {
		fmt.Println(param, "不是回文数")
	}

	//括号匹配
	result3 := task1.IsValidParent("([]){}") // 测试括号匹配
	fmt.Println("括号匹配结果：", result3)
}
