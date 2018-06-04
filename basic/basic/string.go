package main

import (
	"fmt"
	"unicode/utf8"
)

func strings()  {
	s := "在下中二,有个贵干jpg!"  //utf-8

	fmt.Println(s);
	for _,v := range []byte(s){  // uint8
		fmt.Printf("%X ",v)
	}
	fmt.Println()
	for k,ch := range s{   // ch is a rune => uint32
		fmt.Printf("(%d,%X) ",k ,ch)   //便利的过程中 utf8->uint32, k没变化
	}
	fmt.Println()
	for k,ch := range []rune(s){   //  uint32
		fmt.Printf("(%d,%X) ",k ,ch)
	}
	fmt.Println()
	fmt.Println(
		utf8.RuneCountInString(s),
	)

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch ,size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ",ch)
	}
	fmt.Println()
}


func main()  {
	strings()
}
