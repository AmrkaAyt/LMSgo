package main

import "fmt"
import "LMS/Task1"

func main() {
	target := 9
	nums := []int{1, 2, 3, 4, 6, 7, 8}

	fmt.Println(Task1.TwoSum(nums, target))
}
