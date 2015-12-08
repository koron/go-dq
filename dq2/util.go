package dq2

func findRune(t rune, runes []rune) int {
	for i, r := range runes {
		if t == r {
			return i
		}
	}
	return -1
}

func findString(t string, array []string) int {
	for i, s := range array {
		if t == s {
			return i
		}
	}
	return -1
}

func btoi(b bool, n int) int {
	if b {
		return n
	}
	return 0
}
