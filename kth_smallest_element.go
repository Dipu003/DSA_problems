package main

import (
	"fmt"
	"sort"
)

func kthSmallest(arr []int, k int) int {
	sort.Ints(arr)
	return arr[k-1]
}

func main() {
	arr := []int{7, 10, 4, 3, 20, 15}
	k := 3
	result := kthSmallest(arr, k)
	fmt.Println("Kth smallest element:", result)
}
