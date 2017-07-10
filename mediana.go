package main

func Selection(list []int, k int) int {
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
func Median(list []int) int {
	return Selection(list, (len(list)/2)-1)
}
