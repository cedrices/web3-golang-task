package task1

import (
	"strconv"
)

// 返回数组中只出现一次的数字
func ArrayFindSingle(array [11]int) string {
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
