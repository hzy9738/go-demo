package main

import (
	"io/ioutil"
	"fmt"
)

func readfile() {
	const filename = "text.txt"
	content, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%s\n", content)
	}

}

func grade(score int) string {
	g := ""
	switch {
	case score <0 || score>100:
		panic(fmt.Sprintf("Wrong score: %d",score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "c"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"

	}
	return g

}

func main() {
	readfile()
	fmt.Println(
		grade(55),
		grade(65),
		grade(85),
		grade(95),
		grade(100),
		grade(101),
	)
	//grade(1i)
}
