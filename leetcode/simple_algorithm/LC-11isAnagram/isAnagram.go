package main

func isAnagram(s string, t string) bool {
	valueMap := make(map[rune]int)

	for _, c := range s {
		valueMap[c]++
	}
	for _, c := range t {
		valueMap[c]--
	}
	for _, v := range valueMap {
		if v != 0 {
			return false
		}
	}
	return true
}
func main() {

}
