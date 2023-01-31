package main

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		var mid int = low + (high-low)/2
		ans := nums[mid]
		if ans == target {
			return mid
		}
		if target < ans {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
