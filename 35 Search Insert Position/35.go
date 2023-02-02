package main

func searchInsert(nums []int, target int) int {
	var low int = 0
	var high int = len(nums) - 1

	var ans int = 0
	for low <= high {
		var mid int = int(low + (high-low)/2)
		if nums[mid] == target {
			low = mid + 1
			ans = mid + 1
		} else if nums[mid] > target {
			high = mid + 1
		} else {
			return mid
		}
	}
	return ans
}
