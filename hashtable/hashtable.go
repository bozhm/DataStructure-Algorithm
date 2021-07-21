package main

import "fmt"

//哈希表
type HashTable struct {
	arr [7]*TableItem
}

//数据项
type TableItem struct {
	key   int
	value string
	next  *TableItem
}

//初始化
func Init() *HashTable {
	return &HashTable{}
}

//哈希函数
func (receiver *HashTable) HashFunc(key int) int {
	return key % len(receiver.arr)
}

func (receiver *HashTable) Push(key int, value string) {
	index := receiver.HashFunc(key)
	item := receiver.arr[index]
	//封装加入的节点
	tag := &TableItem{key: key, value: value, next: nil}
	if item == nil {
		//该下标没有值，直接写入
		receiver.arr[index] = tag
	} else {
		//判断首节点key是否相等
		if item.key == key {
			item.value = value
			return
		}
		//有值，找到最后一个节点
		for item.next != nil {
			if item.key == key {
				item.value = value
				return
			}
			item = item.next
		}

		//找到最后一个节点,并赋值
		item.next = tag
	}
}

func (receiver *HashTable) Get(key int) interface{} {
	index := receiver.HashFunc(key)
	item := receiver.arr[index]
	if item.key == key {
		return item.value
	}
	for item.next != nil {
		if item.key == key {
			return item.value
		}
		item = item.next
	}
	//判断尾节点key是否相等
	if item.key == key {
		return item.value
	}
	return nil
}

func main() {
	hashTable := Init()
	hashTable.Push(2, "aaaaaa")
	hashTable.Push(2, "bbbbbb")
	hashTable.Push(16, "cccccc")
	fmt.Println(hashTable.Get(2))
	fmt.Println(hashTable.Get(16))
}
