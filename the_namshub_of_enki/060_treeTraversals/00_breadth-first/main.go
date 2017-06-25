// Breadth-first Traversal

/* Breadth-first traversal or BFS is a traversal of a
binary or non-binary tree which prioritises neighbours
over nodes on other levels.

1. Start at the Parent node
2. Starting at its left-most child, identify all first deg children.
3. Continue to identify children, from left to right, sharing
    the same degree.
4. Check for no more chldren, then stop.
*/

/*
Is this the simplest way to traverse a tree?
It looks simple for humans, because you're just reading off
the tree from left to right as you descend levels.
But that structure is a litte complex - to know which
level you're on, you need to already have a whole picture
of the tree.
*/

package main

import "fmt"

func main() {

	fmt.Println("You find a placeholder.")

}
