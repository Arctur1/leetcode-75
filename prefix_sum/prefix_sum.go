package prefix_sum

import "fmt"

func pivotIndex(nums []int) int {
	running_sum := runningSum(nums)
	fmt.Println(running_sum)
	for i := 1; i < len(nums); i++ {
		right, left := 0, 0

		if i > 0 {
			left = running_sum[i-1]
		}
		if i < len(nums) {
			right = running_sum[len(nums)-1] - running_sum[i]
		}

		if left == right {
			return i
		}

	}
	return -1
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}
