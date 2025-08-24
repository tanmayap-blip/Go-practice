package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	val   string
}

func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Print(root.val + " ")
	inOrder(root.right)
}
func preOrder(root *Node) {
	if root == nil {
		return

	}
	fmt.Print(root.val + " ")
	preOrder(root.left)
	preOrder(root.right)
}
func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Print(root.val + " ")

}

func main() {

	root := &Node{val: "+"}
	root.left = &Node{val: "a"}
	root.right = &Node{val: "-"}
	root.right.left = &Node{val: "b"}
	root.right.right = &Node{val: "c"}

	fmt.Println("Inorder traversal")
	inOrder(root)
	fmt.Println()
	fmt.Println("Preorder traversal")
	preOrder(root)
	fmt.Println()
	fmt.Println("Postorder traversal")
	postOrder(root)

}
