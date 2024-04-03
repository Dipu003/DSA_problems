package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var n = len(arr)
	var start = 0
	var end = n - 1

	for start < end {
		var temp = arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
	fmt.Println(arr)
}
