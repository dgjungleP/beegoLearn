package main

func reverseString(s []byte) {
	right := len(s) - 1
	for i := 0; i <= right; i++ {
		s[i], s[right] = s[right], s[i]
		right--
	}

}
func main() {

}
