package main

func containsDuplicate(nums []int) bool {
	var m = make(map[int]bool)
	for _, value := range nums {
		if m[value] == false {
			m[value] = true
		} else {
			return false
		}
	}
	return false
}
