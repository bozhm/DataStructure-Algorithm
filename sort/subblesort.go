package main

import "fmt"

func main() {
	tag := [5]int{9, 8, 3, 7, 2}
	for i := 0; i < len(tag)-1; i++ {
		for j := 0; j < len(tag)-i-1; j++ {
			if tag[j] > tag[j+1] {
				temp := tag[j]
				tag[j] = tag[j+1]
				tag[j+1] = temp
			}
		}

	}
	fmt.Println(tag)

}
