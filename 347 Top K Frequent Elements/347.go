package main

func topKFrequent(nums []int, k int) []int {
	var m = make(map[int]int)

	for _, value := range nums {
		m[value]++
	}

	var bucket = make([][]int, len(nums)+1)
	for key, value := range m {
		bucket[value] = append(bucket[value], key)
	}

	var res []int
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) > 0 {
			res = append(res, bucket[i]...)
		}
		if len(res) == k {
			break
		}
	}
	return res
}
