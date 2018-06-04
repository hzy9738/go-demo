package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}


func print()  {
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println(
		arr[2:5],
		arr[1:6],
		arr[:5],
		arr[2:],
		arr[:],
	)

	s1 := arr[:]
	updateSlice(s1)

	s3 :=arr[2:]

	fmt.Println( s1 ,s3,arr[:])

}

func length()  {
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[3:5]

	fmt.Println(arr,s1,len(s1),cap(s1))
}

func appends()  {
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[3:6]

	s2 := append(s1,10)
	s3 := append(s2,11)
	s4 := append(s3,12)
	s5 := append(s4,13)
	fmt.Println(s1,s2,s3,s4,s5,arr)

}

func makes()  {
	var s []int
	for i:=0 ; i< 100;i++  {
		s = append(s,2*i+1)
	}
	fmt.Println(s)

	s2 := make([]int,16)
	s3 := make([]int,10,32)
	fmt.Println(s2,len(s2),cap(s2))
	fmt.Println(s3,len(s3),cap(s3))
}

func copys()  {
	s1 := []int{2,4,6,8}
	s2 := make([]int,16)
	copy(s2,s1)
	fmt.Println(s2)
}

func deletes()  {
	s1 := []int{2,4,6,8,10,12,14}

	s2 := append(s1[:3],s1[4:]...)  //删除下标为3

	fmt.Println(s2)

	s3 := s1[1:]			//删除首项
	fmt.Println(s3)

	index2 := len(s2)-1
	s4 := s2[:index2]        //删除尾项
	fmt.Println(s4)
}

func main() {
	print()
	length()
	appends()
	makes()
	copys()
	deletes()
}