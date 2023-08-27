package main

import (
	"fmt"
	"math/rand"
	"racingcar/view"
)

type car struct {
	Distance string
}

func main() {
	numCars := view.InputCarNumbers()
	numRounds := view.InputStageNumbers()

	fmt.Println("실행 결과")

	var cars = make([]car, numCars)

	for i := 0; i < numRounds; i++ {
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
