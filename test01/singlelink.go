package main

import "fmt"

type LinkNode struct {
	No       int
	name     string
	NextNode *LinkNode
}

func (receiver *LinkNode) Add(newNode *LinkNode) {
	temp := receiver
	for {
		if temp.NextNode == nil {
			break
		}
		temp = temp.NextNode
	}

	temp.NextNode = newNode
}

func (receiver *LinkNode) Insert(newNode *LinkNode) {
	temp := receiver
	for {
		if temp.NextNode == nil {
			break
			//说明插入temp后面
		} else if temp.NextNode.No > newNode.No {
			break
		} else if temp.NextNode.No == newNode.No {
			fmt.Printf("无法添加")
			return
		}
		temp = temp.NextNode
	}

	newNode.NextNode = temp.NextNode
	temp.NextNode = newNode

}

func (receiver *LinkNode) Show() {

	if receiver.NextNode == nil {
		return
	}
	temp := receiver
	for {
		fmt.Println(temp.NextNode.No, temp.NextNode.name, temp.NextNode)
		if temp.NextNode.NextNode == nil {
			break
		}
		temp = temp.NextNode
	}
}

func main() {
	head := LinkNode{
		0,
		"头节点",
		nil,
	}

	one := &LinkNode{
		No:   1,
		name: "1号",
	}
	two := &LinkNode{
		No:   5,
		name: "5号",
	}
	num := &LinkNode{
		No:   3,
		name: "3号",
	}
	head.Add(one)
	head.Add(two)
	head.Insert(num)
	head.Show()
}
