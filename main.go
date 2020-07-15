package main

import (
	"fmt"
	"strconv"

	"github.com/xxsoul/tree-traversal/tree"
)

func main() {
	tree := new(tree.Tree)
	tree.Degree = 2

	for i := 0; i < 20; i++ {
		tree.AddNode(i)
	}

	testTre(tree)
}

func testTre(tree *tree.Tree) {
	i := 0
	j := 0

	fmt.Print(" |")
	for _, n := range tree.NodesPtrs {
		if i != n.Dynasty {
			fmt.Print("| \r\n |")
			i = n.Dynasty
			j = 0
		}

		if j > 0 {
			fmt.Print("。")
		}

		fmt.Print(strconv.Itoa(n.Data.(int)) + "," + strconv.Itoa(n.Dynasty))
		j++
	}
	fmt.Print("|")
}
