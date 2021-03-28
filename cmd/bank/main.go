package main

import (
	// "bank/pkg/bank/transfer"

	// "bank/pkg/bank/card"
	// "bank/pkg/bank/types"
	"fmt"
)

func main() {
		operations := []int64{2, 24, 4,5,24}
		// sum:=sum(operations)

		fmt.Println(sum(operations))
		fmt.Println(max(operations))
	}
	func sum(operations []int64) int64{
		sum:= int64(0)
		for _, operation := range operations {
			sum += operation
		}
		return sum
	}
	func max(operations []int64)  int64{
		max := operations[0]
		for _, operation := range operations {
			if max <= operation {
	           max = operation
			}
		}
		return max
	// operations := make([]int64, 4)
	// fmt.Println(operations[2])
	// operation := []int64{6,5,6,1,54}
	// fmt.Println(operation[2])
}
