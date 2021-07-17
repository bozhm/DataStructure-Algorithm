package main

import "fmt"

func main() {
	var array [3][6]int
	array[1][2] = 1
	array[2][3] = 2

	for _, v := range array {
		for _, v2 := range v {
			fmt.Print(v2)
		}
		fmt.Println()
	}

	type ValueNode struct {
		row   int
		clo   int
		value int
	}

	var array2 []ValueNode
	array2 = append(array2, ValueNode{
		3,
		6,
		0,
	})

	for i, v := range array {
		for j, v2 := range v {
			if v2 != 0 {
				node := ValueNode{
					i,
					j,
					v2,
				}
				array2 = append(array2, node)
			}
		}
	}

	var ResultArray [3][6]int
	for _, node := range array2 {
		if node.value != 0 {
			ResultArray[node.row][node.clo] = node.value
		}

	}

	fmt.Println("解析结果")
	for _, v := range ResultArray {
		for _, v2 := range v {
			fmt.Print(v2)
		}
		fmt.Println()
	}

}
