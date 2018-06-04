package main

import "fmt"

type treeNode struct {
	value int
	left , right *treeNode
}



//工厂
func createNode(value int) *treeNode  {
	return &treeNode{value: value}
}

func (node treeNode) print()  {
	fmt.Println(node.value," ")
}

func (node *treeNode) setValue(value int)  {
	node.value = value
}

func (node *treeNode) forever() {
	if(node == nil){
		return
	}
	node.left.forever()
	node.print()
	node.right.forever()
}

func main() {
	var root treeNode
	fmt.Println(root)


	root = treeNode{value:4}
	fmt.Println(root)

	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.left.right = new(treeNode)
	root.right.left = createNode(2)

	fmt.Println(root)


	nodes := []treeNode{
		{value:3},
		{},
		{6,nil,nil},
	}
	fmt.Println(nodes)


	root.print()
	root.left.print()

	root.left.setValue(10)
	root.left.print()

	fmt.Println("---")

	root.forever()
}
