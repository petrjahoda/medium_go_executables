package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello world")
	var randomInts []int
	var divisibleInts []int
	for i := 0; i < 10000000; i++ {
		randomInt :=rand.Intn(1000000)
		randomInts = append(randomInts, randomInt)
		if randomInt%10 == 0 {
			divisibleInts = append(divisibleInts, randomInt)
		}
	}
	fmt.Println(len(randomInts))
	fmt.Println(len(divisibleInts))
}
