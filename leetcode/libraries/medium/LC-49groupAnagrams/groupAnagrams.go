package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	resulrMap := map[string][]string{}
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
		resulrMap[string(bytes)] = append(resulrMap[string(bytes)], str)
	}
	result := make([][]string, 0, len(resulrMap))
	for _, v := range resulrMap {
		result = append(result, v)
	}
	return result
}

func main() {

}
