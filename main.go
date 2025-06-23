package main

import (
	"fmt"
)

func main() {

	user1 := map[string]int{
		"ZyWoo is the best player in 2025 year":  1,
		"m0NESY is the best player in 2025 year": 2,
	}
	user2 := map[string]int{
		"ZyWoo1":  1,
		"m0NESY4": 2,
	}
	for k, v := range user1 {
		fmt.Printf("user1 key: %s, value: %d\n", k, v)
	}
	for k, v := range user2 {
		fmt.Printf("user2 key: %s, value: %d\n", k, v)
	}
}
