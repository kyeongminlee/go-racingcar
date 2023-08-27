package main

import (
	"fmt"
	"math/rand"
)

type car struct {
	Distance string
}

func main() {
	var numCars int
	var numRounds int

	fmt.Println("자동차 대수는 몇 대 인가요?")
	fmt.Scanln(&numCars)

	fmt.Println("시도할 회수는 몇 회 인가요?")
	fmt.Scanln(&numRounds)

	fmt.Println("실행 결과")

	var cars = make([]car, numCars)

	for i := 0; i < numRounds; i++ {
		// fmt.Println("===================")
		// fmt.Println("round: ", i)
		for j := 0; j < numCars; j++ {
			randomNumber := rand.Intn(10)
			if randomNumber >= 4 {
				cars[j].Distance = fmt.Sprintf("%v-", cars[j].Distance)
			}
			fmt.Println(cars[j].Distance)
		}
		fmt.Println()
	}
}
