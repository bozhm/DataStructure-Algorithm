package main

import "fmt"

type LinkNode struct {
	No       int
	Name     string
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

func (receiver *LinkNode) Delete(id int) {
	temp := receiver
	for {
		if temp.NextNode == nil {
			fmt.Printf("不存在")
			return
		} else if temp.NextNode.No == id {
			break
		}
		temp = temp.NextNode
	}
	temp.NextNode = temp.NextNode.NextNode
}

func (receiver *LinkNode) Show() {

	if receiver.NextNode == nil {
		return
	}
	temp := receiver
	for {
		fmt.Println(temp.NextNode.No, temp.NextNode.Name, temp.NextNode)
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
		Name: "1号",
	}
	two := &LinkNode{
		No:   5,
		Name: "5号",
	}
	num := &LinkNode{
		No:   3,
		Name: "3号",
	}
	head.Add(one)
	head.Add(two)
	head.Insert(num)
	head.Delete(3)
	head.Show()
}
