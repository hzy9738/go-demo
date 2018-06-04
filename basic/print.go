package basic;

import "fmt"

var (
	aa = 3
	bb = 3
	cc = 3
)


func variableZeroValue() {
	var i int          //默认0
	var n string       //默认""
	fmt.Printf("%d %q\n", i , n);    //%q可以把""打印出来
}

func variableInitiaValue()  {
	var a ,b int
	var c ,d string
	f ,g := 3,"123"
	var e = true
	var h = "黄"
	fmt.Println(a,b,c,d,f,g,e,h)
}



func main() {
	fmt.Println("hello world")
	variableZeroValue();
	variableInitiaValue();
	fmt.Println(aa,bb,cc)
}




/**


hello world

0 ""

0 0   3 123 true 黄

3 3 3



 */
