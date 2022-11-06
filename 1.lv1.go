package main

import (
	"fmt"
)

func main() {
	var n, l, r, x int
	fmt.Scan(&n, &l, &r)
	fmt.Printf("请输入%d个数\n", n)
	var slice1 []int
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		slice1 = append(slice1, x)
	}
	slice2 := slice1[l : r+1]
	for j := 0; j < len(slice2)-1; j++ {
		for k := 0; k < len(slice2)-1-j; k++ {
			if slice2[k] > slice2[k+1] {
				slice2[k], slice2[k+1] = slice2[k+1], slice2[k]
			}
		}
	}
	slice3 := slice1[0:l]
	for _, value1 := range slice3 {
		fmt.Printf("%d ", value1)
	}
	for _, value2 := range slice2 {
		fmt.Printf("%d ", value2)
	}

}
