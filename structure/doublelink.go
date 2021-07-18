package main

import "fmt"

type DoubleLinkNode struct {
	No       int
	Name     string
	PreNode  *DoubleLinkNode
	NextNode *DoubleLinkNode
}

func (receiver *DoubleLinkNode) Add(newNode *DoubleLinkNode) {
	temp := receiver
	for {
		if temp.NextNode == nil {
			break
		}
		temp = temp.NextNode
	}

	temp.NextNode = newNode
	newNode.PreNode = temp
}

func (receiver *DoubleLinkNode) Insert(newNode *DoubleLinkNode) {
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
	temp.NextNode.PreNode = newNode

	temp.NextNode = newNode
	newNode.PreNode = temp
}

func (receiver *DoubleLinkNode) Delete(id int) {
	temp := receiver
	for {
		if temp.NextNode == nil {
			fmt.Printf("不存在")
			return
		} else if temp.No == id {
			break
		}
		temp = temp.NextNode
	}
	temp.NextNode.PreNode = temp.PreNode
	temp.PreNode.NextNode = temp.NextNode
}

func (receiver *DoubleLinkNode) Show() {

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
func (receiver *DoubleLinkNode) BackShow() {
	temp := receiver
	if temp.NextNode == nil {
		return
	}
	for {
		if temp.NextNode == nil {
			break
		}
		temp = temp.NextNode
	}

	for {
		fmt.Println(temp.No, temp.Name, temp.PreNode)
		if temp.PreNode == nil {
			break
		} else if temp.PreNode.PreNode == nil {
			break
		}
		temp = temp.PreNode
	}
}

func main() {
	head := DoubleLinkNode{
		0,
		"头节点",
		nil,
		nil,
	}

	one := &DoubleLinkNode{
		No:   1,
		Name: "1号",
	}
	two := &DoubleLinkNode{
		No:   5,
		Name: "5号",
	}
	num := &DoubleLinkNode{
		No:   3,
		Name: "3号",
	}
	head.Add(one)
	head.Add(two)
	head.Insert(num)
	head.Delete(1)
	head.Show()
	//head.BackShow()
}
