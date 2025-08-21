package main

import "fmt"

func main() {
	//两数之和
	nums := []int{2, 11, 7, 15}
	target := 9
	result := twoSum(nums, target)
	for _, value := range result {
		fmt.Printf("nums数组下标[%d]对应值   \n", value)
	}
	fmt.Println("两数之和：", target)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
