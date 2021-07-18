package main

import "fmt"

func main() {
	tag := [5]int{9, 8, 3, 7, 2}
	for i := range tag {
		preIndex := i - 1
		current := tag[i]
		for preIndex >= 0 && tag[preIndex] > current {
			tag[preIndex+1] = tag[preIndex]
			preIndex -= 1
		}
		tag[preIndex+1] = current
	}
	fmt.Println(tag)
}
