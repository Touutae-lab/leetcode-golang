package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	hash := make(map[string][]string)

	for _, str := range strs {
		sorted := sortString(str)
		hash[sorted] = append(hash[sorted], str)
	}
	return hashToSlice(hash)
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func hashToSlice(hash map[string][]string) [][]string {
	var result [][]string
	for _, v := range hash {
		result = append(result, v)
	}
	return result
}
