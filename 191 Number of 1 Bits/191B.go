package main

func hammingWeight(num uint32) int {
	//return bits.OnesCount32(num)

	var count int
	for num > 0 {
		count += int(num & 1)
		num = num >> 1
	}
	return count
}
