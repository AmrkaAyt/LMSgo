func twoSum(nums []int, target int) []int {
	for i, r := range nums {
		for j := i + 1; j < len(nums); j++ {
			if r+nums[j] == target && i < j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
