package main

func isAnagram(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}
	ctr := make([]int, 26)
	for i := range s {
		ctr[s[i]-'a']++
		ctr[t[i]-'a']--
	}

	for _, c := range ctr {
		if c != 0 {
			return false
		}
	}
	return true
}
