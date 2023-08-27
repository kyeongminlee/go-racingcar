package view

import (
	"fmt"
	"racingcar/domain"
)

var MOVEMENT = "-"

func PrintResultMention() {
	fmt.Println("======= 실행 결과 =======")
}

func PrintCarsStatus(cars domain.Cars) {
	for _, car := range cars {
		printCarStatus(car.GetCurrentPosition())
	}
}

func printCarStatus(carPosition int) {
	for i := 0; i < carPosition; i++ {
		fmt.Print(MOVEMENT)
	}
	fmt.Println()
}
