package main

import "fmt"

func main() {
	arr := []int{254, 430, 612, 72, 85, 1010}
	fmt.Println(FindOutlier(arr))
}

func FindOutlier(integers []int) int {
	var evenSlice []int
	var oddSlice []int

	for _, digit := range integers {
		if digit%2 == 0 {
			evenSlice = append(evenSlice, digit)
		} else {
			oddSlice = append(oddSlice, digit)
		}
	}
	if len(evenSlice) == 1 {
		return evenSlice[0]
	} else {
		return oddSlice[0]
	}
}
