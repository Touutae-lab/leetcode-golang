package main

func groupAnagrams(strs []string) [][]string {
	// we store the frequency of each letter in a string as a key
	// by convert the string to a byte array we don't have to check array
	freqs := make(map[[26]byte][]string, len(strs))

	for _, s := range strs {
		g := [26]byte{}

		for _, c := range s {
			g[c-'a']++
		}

		freqs[g] = append(freqs[g], s)
	}

	result := [][]string{}

	for _, v := range freqs {
		result = append(result, v)
	}

	return result
}
