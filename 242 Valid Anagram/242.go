package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var m = make(map[rune]int)

	for _, value := range s {
		m[value]++
	}

	for _, value := range t {
		m[value]--
	}

	for _, value := range m {
		if value != 0 {
			return false
		}
	}

	return true
}
