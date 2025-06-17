package main

import "fmt"

func main() {

	users := map[string]map[string]int{
		"user1": {
			"ZyWoo1":  1,
			"m0NESY2": 2,
			"Donk3":   3,
			"Simleo4": 4,
			"nIKo5":   5,
		},
		"user2": {
			"ZyWoo1":  1,
			"m0NESY2": 2,
			"Donk3":   3,
			"Simleo4": 4,
			"nIKo6":   6,
		},
	}
	for k, v := range users["user1"] {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
	for k, v := range users["user2"] {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
}
