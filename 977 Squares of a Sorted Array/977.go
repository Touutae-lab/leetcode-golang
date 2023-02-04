package main

func pow(x int) int {
	return x * x
}

func sortedSquares(nums []int) []int {
	var length int = len(nums) - 1
	var start int = 0
	var end int = length

	var answer = make([]int, length+1)

	for start <= end && length > -1 {
		if pow(nums[start]) > pow(nums[end]) {
			answer[length] = pow(nums[start])
			length -= 1
			start += 1

		} else {
			answer[length] = pow(nums[end])
			end -= 1
			length -= 1
		}
	}
	return answer
}
