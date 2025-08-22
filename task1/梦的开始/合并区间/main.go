package main

import (
	"fmt"
)

func main() {

	//合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {17, 20}}
	intervals = mergeIntervals(intervals)
	fmt.Println("合并区间结果：", intervals)
}

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}

	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		if current[0] <= last[1] { // 有重叠
			last[1] = max(last[1], current[1])
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
