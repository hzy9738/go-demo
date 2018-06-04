package main

import "fmt"


// 寻找最长不含有重复字符的字串长度 （不支持中文）
func norepeatingSubStr(s string) int {
	last0ccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		lastI, ok := last0ccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		last0ccurred[ch] = i
	}

	return maxLength

}


// 支持中文

func norepeatingSubStrCN(s string) int  {
	last0ccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		lastI, ok := last0ccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		last0ccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(
		norepeatingSubStr("12312312"),
		norepeatingSubStr("abcabcabc"),
		norepeatingSubStr("iii"),
		norepeatingSubStr(""),
		norepeatingSubStr("v"),
		norepeatingSubStr("hahah"),
	)
	fmt.Println(
		norepeatingSubStrCN("hzy沐青1之枫123"),
	)
}
