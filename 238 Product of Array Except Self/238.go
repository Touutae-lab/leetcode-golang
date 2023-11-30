package main

func productExceptSelf(nums []int) []int {
	var prefix = make([]int, len(nums))
	var suffix = make([]int, len(nums))

	prefix[0] = nums[0]
	suffix[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i]
		suffix[len(nums)-1-i] = suffix[len(nums)-i] * nums[len(nums)-1-i]
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums[i] = suffix[i+1]
		} else if i == len(nums)-1 {
			nums[i] = prefix[i-1]
		} else {
			nums[i] = prefix[i-1] * suffix[i+1]
		}
	}

	return nums
}
