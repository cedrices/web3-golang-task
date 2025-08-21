package main

import (
	"fmt"
	"strconv"
)

func main() {
	//只出现一次的数字
	arrays := [11]int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5}
	result1 := arrayFindSingle(arrays)
	fmt.Println("只出现一次的数字是：", result1)
}

// 返回数组中只出现一次的数字
func arrayFindSingle(array [11]int) string {
	var counts = make(map[int]int, 11)
	for _, v := range array {
		var count = counts[v]
		count++
		counts[v] = count
	}
	for key, value := range counts {
		if value == 1 {
			return strconv.FormatInt(int64(key), 10)
		}
	}
	return ""
}
