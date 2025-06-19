package main

import (
	"fmt"
)

func main() {

	user1 := map[string]int{
		"ZyWoo is the best player in 2025 year":  1,
		"m0NESY is the best player in 2025 year": 2,
		"Donk is the best player in 2025 year":   3,
		"Simpleo4":                               4,
		"nIKo5":                                  5,
	}
	user2 := map[string]int{
		"ZyWoo1":   1,
		"m0NESY2":  2,
		"Donk3":    3,
		"Simpleo4": 4,
		"nIKo5":    5,
	}
	for k, v := range user1 {
		fmt.Printf("user1 key: %s, value: %d\n", k, v)
	}
	for k, v := range user2 {
		fmt.Printf("user2 key: %s, value: %d\n", k, v)
	}
}
