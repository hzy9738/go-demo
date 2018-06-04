package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(n int) string  {
	result := ""
	for ; n > 0 ; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return  result
}

func printfile(filename string)  {
	file,error := os.Open(filename)
	if error != nil {
		panic(error)
	}
	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever()  {
	//死循环
	for {
		fmt.Println("abc")
	}
	
}


func main()  {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1011 ->1101
		convertToBin(520),

	)
	printfile("text.txt")
	forever();
}