package main

import "fmt"

/**
循环链表
*/

type AnnularLink struct {
	No       int
	name     string
	NextNode *AnnularLink
}

func Add(headLink *AnnularLink, newLink *AnnularLink) {
	//第一次
	if headLink.NextNode == nil {
		headLink.No = newLink.No
		headLink.name = newLink.name
		headLink.NextNode = headLink
		return
	}
	temp := headLink
	for {
		if temp.NextNode == headLink {
			break
		}
		temp = temp.NextNode
	}
	newLink.NextNode = temp.NextNode
	temp.NextNode = newLink
}

func Show(headLink *AnnularLink) {
	temp := headLink
	for {
		fmt.Println(temp.No, temp.name, temp.NextNode)
		if temp.NextNode == headLink {
			return
		}
		temp = temp.NextNode
	}

}

func main() {

	headlink := &AnnularLink{}
	one := &AnnularLink{
		No:   1,
		name: "张三",
	}
	two := &AnnularLink{
		No:   2,
		name: "李四",
	}
	tree := &AnnularLink{
		No:   3,
		name: "王五",
	}
	Add(headlink, one)
	Add(headlink, two)
	Add(headlink, tree)
	Show(headlink)
}
