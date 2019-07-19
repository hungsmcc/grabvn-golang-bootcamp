package main

type Tree struct {
	Left  *Tree
	Value string
	Right *Tree
}

//1 + 2 * 4

func insert(t *Tree, value string) *Tree {

	switch value {
	case "+", "-":
		if t == nil {
			//new leaf node
			return &Tree{nil, value, nil}
		} else {
			return &Tree{t, value, nil}
		}
	case "*", "/":
		if t == nil {
			//new leaf node
			return &Tree{nil, value, nil}
		} else {
			//create new current t.right with left = old right
			//newRight := &Tree{t.Right, value, nil}
			//fmt.Println("before ", t.Right.Value)

			t.Right = &Tree{t.Right, value, nil}
			//t.Right = insert(&Tree{t.Right, value, nil}, value)

			//fmt.Println("after ", t.Right.Value)
			//fmt.Println("after ", t.Right.Left.Value)
			//fmt.Println("after ", t.Right.Right.Value)
			return t
		}

	default:
		if t == nil {
			//new leaf node
			return &Tree{nil, value, nil}
		} else {
			//always deepest right
			//t.Right = &Tree{nil, value, nil}
			t.Right = insert(t.Right, value)
			return t
		}
	}
}

func TraversePostOrder(t *Tree) {
	if t == nil {
		return
	}

	// first recursive on left subtree
	TraversePostOrder(t.Left)

	// then recursive on right subtree
	TraversePostOrder(t.Right)

	//cal node value
	CalculateNodeByStack(t.Value)
}

func New(a []string) *Tree {
	var t *Tree
	for _, v := range a {
		t = insert(t, v)
	}
	return t
}
