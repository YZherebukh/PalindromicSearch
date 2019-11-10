package main

import (
	"flag"
	"fmt"
)

func main() {
	numbPtr := flag.Int("numb", 0, "an int")

	flag.Parse()

	if *numbPtr == 0 {
		fmt.Println("Please, enter number grater than 0 ")
		return
	}

	h, l := getUpperLine(*numbPtr)
	highest := h * h

	for i := highest; i > l * l; i-- {
		if checkPalidrom(i) {
			n1, n2, n3 := checkIfDevided(i, h, l)
			if n1 != 0 {
				fmt.Printf("%d * %d = %d \n", n1, n2, n3)
				break
			}
		}
	}
}

func checkPalidrom(p int) bool {
	var original, digit int
	original = p
	numbers := make([]int, 0)
	for {
		digit = p % 10
		numbers = append(numbers, digit)
		p = p / 10
		if p == 0 {
			break
		}
	}

	var flippedP int
	for i, v := range numbers {
		i++
		for t := 0; t < len(numbers)-i; t++ {
			v = v * 10
		}
		flippedP = flippedP + v
	}

	return original == flippedP
}

func getUpperLine(n int) (highest int, lowest int) {
	highest = 1
	lowest = 1
	for i := 0; i < n; i++ {
		highest = highest * 10
	}
	lowest = highest / 10
	highest = highest - 1
	return

}

func checkIfDevided(n, h, l int) (int, int, int) {
	if h < l {
		return 0,0,0
	}
	
	n1 := n / h
	if n1*h == n && n1 > l && n1 < h {
		return n1, h, n
	}

	return checkIfDevided(n, h-1, l)
}