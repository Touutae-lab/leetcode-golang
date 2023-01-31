package main

func firstBadVersion(n int) int {
	var low int = 1
	var high int = n
	for true {
		var mid int = low + (high-low)/2
		if isBadVersion(mid) {
			if isBadVersion(mid - 1) {
				high = mid - 1
			} else {
				return mid
			}
		} else {
			if isBadVersion(mid + 1) {
				return mid + 1
			} else {
				low = mid + 2
			}
		}
	}
	return -1
}
