package main

func rotate(nums []int, k int) {
	k %= len(nums)
	if k == 0 {
		return
	}

	reverse := func(nums []int) {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	pivot := len(nums) - k
	reverse(nums[pivot:])
	reverse(nums[:pivot])
	reverse(nums)
}
