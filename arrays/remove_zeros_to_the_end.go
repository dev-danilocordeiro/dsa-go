package main

import "fmt"

func maine()  {
	fmt.Println(moveZeros([]int{0,0,0,0,0,1,0,2,0,0,3,4,0,5,0,0,0,0,0}))
}

func moveZeros(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 && arr[j] == 0 {
			arr[i], arr[j] = arr[j], arr[i]
		}

		if arr[j] != 0 {
			j++
		}
	}

	return arr
}
