package main

// we can easily do this by slice array
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	var slice1 []int = make([]int, 0)
	var pivot int = len(nums) - k
	slice1 = append(nums, nums[:pivot]...)
	copy(nums, slice1[pivot:])
}
