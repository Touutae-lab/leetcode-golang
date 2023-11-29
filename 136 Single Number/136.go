package main

func singleNumber(nums []int) int {
	var m int = 0

	for _, value := range nums {
		m ^= value
	}
	return m
}
