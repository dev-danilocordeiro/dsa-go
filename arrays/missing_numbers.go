package main

import "fmt"


func maint()  {
	fmt.Println(missingNumbers([]int{3,0,1,5,4,7,8}))
}

func missingNumbers(nums []int) int {
	arrLength := len(nums)
	sum := arrLength * (arrLength + 1) / 2

	for _, value := range nums {
		sum -= value
	}

	return sum
}