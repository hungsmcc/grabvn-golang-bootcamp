package main

import "fmt"

func main() {
	a := []string{"1", "+", "2", "*", "4", "/", "4", "-", "5", "*", "2"}
	t1 := New(a)

	//fmt.Println(t1.Value)
	//fmt.Println(t1.Left.Value)
	//fmt.Println(t1.Right.Value)
	//
	//fmt.Println(t1.Right.Left.Value)
	//fmt.Println(t1.Right.Right.Value)

	TraversePostOrder(t1)

	fmt.Println()
}
