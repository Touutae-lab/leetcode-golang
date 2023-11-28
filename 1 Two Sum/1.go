package main

func twoSum(nums []int, target int) []int {

	hash := make(map[int]int)

	for i, v := range nums {
		if j, ok := hash[target-v]; ok {
			return []int{j, i}
		}
		hash[v] = i
	}
	return []int{-1}
}
