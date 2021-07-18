package main

import "fmt"

func main() {
	tag := [5]int{9, 8, 3, 7, 2}
	for i := 0; i < len(tag)-1; i++ {
		min := i
		for j := min + 1; j < len(tag); j++ {
			if tag[min] > tag[j] {
				min = tag[j]
			}
		}
		tag[i], tag[min] = tag[min], tag[i]
	}
	fmt.Println(tag)

}
