package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	total := 0
	x, err := strconv.Atoi(os.Args[1])
	if err == nil {
		for {
			x = cal(x)
			if x <= 0 {
				break
			}
			total += x
		}
	}
	fmt.Println(total)
}

func cal(x int) int {
	if x >= 0 {
		return x/3 - 2
	} else {
		return 0
	}
}
