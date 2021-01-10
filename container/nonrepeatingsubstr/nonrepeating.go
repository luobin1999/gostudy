package substr

func nonRepeatingSubstr(str string) int {
	if len(str) == 0 {
		return 0
	}
	chrr := []rune(str)
	start, maxlen := 0, 0
	record := map[rune]int{}
	for i, ch := range chrr {
		index, ok := record[ch]
		if ok && index >= start {
			start = index + 1
		}
		if i-start+1 > maxlen {
			maxlen = i - start + 1
		}
		record[ch] = i
	}
	return maxlen
}
