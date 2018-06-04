package main

import "fmt"

func array()  {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{23,4,4124,1}
	var grid [4][5]bool

	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)
}

func foreach(arr [3]int)  {
	for k,v := range arr{
		fmt.Println(k ,v)
	}
}


func main() {
	array()
	foreach([3]int{1,3,5} )
}
