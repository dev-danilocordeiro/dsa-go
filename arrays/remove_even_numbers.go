package main

import "fmt"

func mainw()  {
	fmt.Println(removeEvenNumbersFromArray([]int{1,2,3,4,5}))
}

func removeEvenNumbersFromArray(arr []int) (oddNumbers []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			oddNumbers = append(oddNumbers, arr[i])
		}
	}

	return oddNumbers
}