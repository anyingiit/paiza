package main

import "fmt"

func main() {
	var ints [][]int
	for i := 0; ; i++ {
		ints = append(ints, []int{})
		var temp int
		for j := 0; ; j++ {
			fmt.Scan(&temp)
			if temp == -1 || temp == -2 {
				break
			}
			ints[i] = append(ints[i], temp)
		}
		if temp == -2 {
			break
		}
	}
	fmt.Println(ints)
}
