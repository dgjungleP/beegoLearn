package main

func firstUniqChar(s string) int {
	temp := make(map[rune]int)
	for i, v := range s {
		if _, ok := temp[v]; !ok {
			temp[v] = i
		} else {
			temp[v] = -1
		}
	}
	for i := 0; i < len(s); i++ {
		if num, ok := temp[rune(s[i])]; ok && num != -1 {
			return i
		}
	}
	return -1
}

func main() {

}
