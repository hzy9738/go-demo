package main

import "fmt"

func maps()  {
	m := map[string]string{
		"name" : "hzy",
		"site" : "18",
		"height" : "170",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m,m2,m3)

	for k,v := range  m  {
		fmt.Println(k + "--"+  v)
	}
}


func exists(){
	m := map[string]string{
		"name" : "hzy",
		"site" : "18",
		"height" : "170",
	}
	s1 , ok := m["name"];

	fmt.Println(s1,ok)

	s2 , ok := m["sex"];
	fmt.Println(s2,ok)


	if s3 , ok := m["weight"]; ok {
		fmt.Println(s3)
	}else{
		fmt.Println("weight is not exist")
	}
}

func main(){
	maps()
	exists()
}