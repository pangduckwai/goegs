package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tree a binary tree
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func treeNew(k int) *Tree {
	a := rand.Perm(10)
	fmt.Println(a) // TEMP
	t := &Tree{nil, (a[0] + 1) * k, nil}

	for i := 1; i < len(a); i++ {
		populate(t, (a[i]+1)*k)
	}

	return t
}
func populate(t *Tree, v int) {
	if v < t.Value {
		if t.Left != nil {
			populate(t.Left, v)
		} else {
			t.Left = &Tree{nil, v, nil}
		}
	} else if v > t.Value {
		if t.Right != nil {
			populate(t.Right, v)
		} else {
			t.Right = &Tree{nil, v, nil}
		}
	}
}

// Walk walk a b-tree
func Walk(t *Tree, ch chan int) { // *tree.Tree
	move(t, ch)
	close(ch)
}
func move(t *Tree, ch chan int) { // *tree.Tree
	if t.Left != nil {
		move(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		move(t.Right, ch)
	}
}

// Same compare 2 b-trees
func Same(t1, t2 *Tree) bool { // *tree.Tree
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	isSame := true
	for {
		v1, b1 := <-c1
		v2, b2 := <-c2

		if (!b1 && b2) || (b1 && !b2) {
			fmt.Println("One channel closed early") // TEMP
			isSame = false
			break
		} else if !b1 && !b2 {
			break
		}

		if v1 != v2 {
			isSame = false
			break
		}
	}

	return isSame
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// c := make(chan int)
	// go Walk(treeNew(1), c) // tree.New(1)
	// for i := range c {
	// 	fmt.Println(i)
	// }
	fmt.Println("Expected the same :", Same(treeNew(1), treeNew(1)))
	fmt.Println()
	fmt.Println("Expected different:", Same(treeNew(1), treeNew(2)))
}
