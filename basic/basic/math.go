package main

import (

	"math/cmplx"
	"fmt"
	"math"
)


func euler()  {
	c := 3 + 4i
	// 欧拉公式
	euler := cmplx.Exp(1i * math.Pi) + 1
	//euler := cmplx.Pow( math.E, 1i * math.Pi ) + 1;
	fmt.Println(cmplx.Abs(c)  , euler)
}


//强制类型转换
func taiangle()  {
	var a ,b int = 3 ,4
	var c int
	c = int( math.Sqrt(  float64(a * a + b * b ) ));
	fmt.Println(c)
}

const  filename  = "file.text"

func consts()  {
	const  a,b  = 3,4  //常量可以不规定类型
	var c int
	c =  int( math.Sqrt( a * a + b * b  ));
	fmt.Println(filename,c)
}


func enums()  {
	// iota 决定常量内是自增的
	const (
		js     =   iota
		python
		golang
		php
	)

	const (
		b = 1 << (10 + iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(js , python , golang ,php )
	fmt.Println(b , kb, mb, gb, tb, pb)
}



func main()  {
	euler()
	taiangle()
	consts()
	enums()
}


/**


5 (0+1.2246467991473515e-16i)
5



 */