package main

import (
	"fmt"
	"strconv"
)

//stack= pop and push(append)
func pop(stack []int) (int, []int) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

//global
var stack []int
var a, b int

func CalculateNodeByStack(c string) {

	switch {

	case c == "+":
		b, stack = pop(stack)
		a, stack = pop(stack)
		stack = append(stack, a+b)

	case c == "-":
		b, stack = pop(stack)
		a, stack = pop(stack)
		stack = append(stack, a-b)

	case c == "*":
		b, stack = pop(stack)
		a, stack = pop(stack)
		stack = append(stack, a*b)

	case c == "/":
		b, stack = pop(stack)
		a, stack = pop(stack)
		stack = append(stack, a/b)

	default:
		num, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println(err)
		} else {
			stack = append(stack, num)
		}
	}

	//fmt.Println("Stack items:", stack)
}
