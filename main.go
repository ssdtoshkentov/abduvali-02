package main

import "fmt"

func main() {
	user1 := map[string]int{
		"ZyWoo1":  1,
		"m0NESY2": 2,
		"Donk3":   3,
		"Abduvali's new player cs2 and he is 13 years old": 1,
		"batyor5": 5,
	}
	user2 := map[string]int{
		"ZyWoo1":  1,
		"m0NESY2": 2,
		"Donk3":   3,
		"Abduvali's new player cs2 and he is 13 years old": 1,
		"baxtyor5": 5,
	}
	for k, v := range user1 {
		fmt.Printf("user1 key: %s, value: %d\n", k, v)
	}
	for k, v := range user2 {
		fmt.Printf("user2 key: %s, value: %d\n", k, v)
	}
}
