package main

// move left if number less than tartget
// move rifht if it larger
func twoSum(numbers []int, target int) []int {
	first := 0
	second := len(numbers) - 1
	for first < second {
		if numbers[first]+numbers[second] == target {
			break
		} else if numbers[first]+numbers[second] < target {
			first += 1
		} else {
			second -= 1
		}
	}
	return []int{first + 1, second + 1}
}
