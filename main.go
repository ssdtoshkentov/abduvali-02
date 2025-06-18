package main

import "fmt"

func main() {
	user1 := map[string]int{
		"ZyWoo1":  1,
		"m0NESY2": 2,
		"Donk3":   3,
		"Simleo4": 4,
		"batyor5": 5,
	}
	user2 := map[string]int{
		"ZyWoo1":   1,
		"m0NESY2":  2,
		"Donk3":    3,
		"Simleo4":  4,
		"baxtyor5": 5,
	}
	for k, v := range user1 {
		fmt.Printf("user1 key: %s, value: %d\n", k, v)
	}
	for k, v := range user2 {
		fmt.Printf("user2 key: %s, value: %d\n", k, v)
	}
}
