package main

import (
	"fmt"
)

func main() {

	//删除有序数组中的重复项
	nums := []int{1, 1, 2, 3, 3, 4}
	result, length := removeDuplicates(nums)
	for _, num := range result {
		print(num, " ")
	}
	fmt.Println("删除重复项后的数组：", result)
	fmt.Println("新数组的长度：", length)
}

func removeDuplicates(nums []int) ([]int, int) {
	if len(nums) == 0 {
		return nums, 0
	}
	if len(nums) == 1 {
		return nums, 1
	}

	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}

	}
	return nums[:slow], slow
}
