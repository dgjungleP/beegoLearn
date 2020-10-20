package main

func firstUniqChar(s string) byte {
	valueMap := make(map[byte]int)

	byteArr := []byte(s)

	for i := 0; i < len(byteArr); i++ {
		valueMap[byteArr[i]] = valueMap[byteArr[i]] + 1
	}
	for i := 0; i < len(byteArr); i++ {
		if value, _ := valueMap[byteArr[i]]; value == 1 {
			return byteArr[i]
		}
	}
	return ' '
}

func main() {
	firstUniqChar("absnsban")
}
