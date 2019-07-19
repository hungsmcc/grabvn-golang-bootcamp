package main

import (
	"fmt"
)

func main() {

	//scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	//for scanner.Scan() {
	//text := scanner.Text()

	//test a fix array
	a := []string{"1", "+", "2", "*", "4", "/", "4", "-", "5", "*", "2"}
	//a := []string{"6", "*", "2"}

	//a := strings.Fields(text)
	//fmt.Println(text)

	//array to tree
	t1 := New(a)

	//test tree
	//fmt.Println(t1.Value)
	//fmt.Println(t1.Left.Value)
	//fmt.Println(t1.Right.Value)
	//
	//fmt.Println(t1.Right.Left.Value)
	//fmt.Println(t1.Right.Right.Value)

	//traverse tree
	TraversePostOrder(t1)

	//get result
	lastItem, _ := pop(stack)

	fmt.Println(a, " = ", lastItem)

	fmt.Print("> ")
	//}
}
