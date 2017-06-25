/*
Applicable only to binary trees.
° Start at the top.
° If there is a new left child, check for and name children
° If there are no new children, return to previous level.
° If there is a new right child, check for and name children

... this is root, leftChild, rightChild

This can be simplified:
Q1: If there is a new left child, go to the child, and ask this again, otherwise go to Q2.
Q2: If there is a new right child, go to the child and ask Q1, otherwise go up a level.

*/

package main

import "fmt"

func main() {

	fmt.Println("You find a placeholder.")
}
