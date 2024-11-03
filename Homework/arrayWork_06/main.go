package main

import "fmt"

// goland教程中数组作业

func main() {
	numberArray := [...]int{1, 3, 5, 7, 8}
	sum := 0

	for i := 0; i < len(numberArray); i++ {
		sum += numberArray[i]
	}
	fmt.Println("总数是", sum)

	for i := 0; i < len(numberArray); i++ {
		for j := i; j < len(numberArray); j++ {
			if numberArray[i]+numberArray[j] == 8 {
				fmt.Println("据计算，角标为", i, "和", j, "时，两数相加为8")
			}
		}
	}

}
