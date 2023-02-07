package main

func reverseWords(s string) string {
	answer := []byte(s)
	lastSpaceIndex := -1
	for i := 0; i <= len(answer); i++ {
		if i == len(answer) || answer[i] == ' ' {
			start, end := lastSpaceIndex+1, i-1
			for i, j := start, end; i < j; i, j = i+1, j-1 {
				answer[i], answer[j] = answer[j], answer[i]
			}
			lastSpaceIndex = i
		}
	}
	return string(answer)
}
