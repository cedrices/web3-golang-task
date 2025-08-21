package main

import (
	"fmt"

	"github.com/cedrices/web3-golang-task/task1"
)

func main() {
	//只出现一次的数字
	// arrays := [11]int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5}
	// result := task1.ArrayFindSingle(arrays)
	// fmt.Println(result)
	// fmt.Println("Hello World!")
	var result = task1.VerifyPalindrome(-121)
	fmt.Println(result)
}
