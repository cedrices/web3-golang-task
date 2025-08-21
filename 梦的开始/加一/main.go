package main

import (
	"fmt"
)

func main() {

	//加一
	arr := []int{1, 2, 9}
	result5 := addOne(arr)
	fmt.Println("加一后的结果是：", result5)
}

// 加一
func addOne(arr []int) []int {
	if len(arr) == 0 {
		return []int{1}
	}

	carry := 1
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] += carry
		if arr[i] == 10 {
			arr[i] = 0
			carry = 1
		} else {
			carry = 0
			break
		}
	}

	if carry == 1 {
		arr = append([]int{1}, arr...)
	}
	return arr
}
