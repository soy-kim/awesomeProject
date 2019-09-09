package main

import "fmt"

func main() {
	for i := 99; i > 0; i-- {
		if i > 1 {
			fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", i, i)
			if i-1 == 1 {
				fmt.Printf("Take one down, pass it around, %d bottle of beer on the wall\n", i-1)
			} else {
				fmt.Printf("Take one down, pass it around, %d bottles of beer on the wall\n", i-1)
			}
		} else {
			fmt.Printf("%d bottle of beer on the wall, %d bottle of beer\n", i, i)
			fmt.Printf("Take one down, pass it around, No more bottles of beer on the wall\n")
		}
	}
}
