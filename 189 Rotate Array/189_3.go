package main

func rotate(nums []int, k int) {
	n := len(nums)
	if k %= n; k == 0 {
		return
	}
	for i, count := 0, 0; count < n; i++ {
		j := i
		curr := nums[j]
		for {
			j = (j + k) % n
			next := nums[j]
			nums[j] = curr
			curr = next
			count++
			if j == i {
				break
			}
		}
	}
}
