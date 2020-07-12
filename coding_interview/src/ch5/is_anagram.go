package ch5string

// 242

func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	stat := make(map[rune]int)
	for _, ch := range str1 {
		stat[ch]++
	}
	for _, ch := range str2 {
		stat[ch]--
		if stat[ch] < 0 {
			return false
		}
	}

	return true
}
