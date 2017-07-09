package main

import (
	"fmt"
	// "net/http"
	// "reflect"
	//  "os"
)

func selection(list []int, k int) int {
	n := len(list)
	for i := 0; i <= k; i++ {
		minIndex := i
		minValue := list[i]
		for j := i; j < n; j++ {
			if list[j] < minValue {
				minIndex = j
				minValue = list[j]
				list[i], list[minIndex] = list[minIndex], list[i]
			}
		}
	}
	return list[k]
}
func mainnn() {
	list := []int{2, 2, 3, 4, 45, 62, 123, 4, 435, 45, 7, 56345, 2345, 202, 35446, 65434534, 45, 2, 64, 332, 433, 5, 7, 11, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	//list :=[]int{1,2,3,4,5}
	fmt.Println(selection(list, 4))
	fmt.Println(list)
}
