package main

import "fmt"

type Node struct {
	Name      string
	LeftNode  *Node
	RightNode *Node
}

// PreorderTraversal 前序遍历  根-> 左-> 右
func PreorderTraversal(node *Node) {
	if node != nil {
		fmt.Println(node.Name)
		PreorderTraversal(node.LeftNode)
		PreorderTraversal(node.RightNode)
	}

}

// InorderTraversal 中序遍历   左-> 根-> 右
func InorderTraversal(node *Node) {
	if node != nil {
		PreorderTraversal(node.LeftNode)
		fmt.Println(node.Name)
		PreorderTraversal(node.RightNode)

	}

}

// PostorderTraversal 后续遍历 左-> 右-> 根
func PostorderTraversal(node *Node) {
	if node != nil {
		PreorderTraversal(node.LeftNode)
		PreorderTraversal(node.RightNode)
		fmt.Println(node.Name)
	}

}

func main() {
	rootNode := &Node{Name: "根节点"}

	l1 := &Node{Name: "l1"}
	r1 := &Node{Name: "r1"}

	l2 := &Node{Name: "l2"}
	l3 := &Node{Name: "l3"}
	l1.LeftNode = l2
	l1.RightNode = l3

	r2 := &Node{Name: "r2"}
	r1.LeftNode = r2

	rootNode.LeftNode = l1
	rootNode.RightNode = r1
	PreorderTraversal(rootNode)
	fmt.Println("-----------------------------")
	InorderTraversal(rootNode)
	fmt.Println("-----------------------------")
	PostorderTraversal(rootNode)
}
