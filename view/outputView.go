package view

import (
	"fmt"
	"racingcar/domain"
)

var MOVEMENT = "-"

func PrintResultMention() {
	fmt.Println("======= 실행 결과 =======")
}

func PrintCarsStatus(cars domain.Cars, stages domain.Stages) {
	fmt.Printf("remain stages: %v \n", stages.CurrentStage()+1)
	for _, car := range cars {
		fmt.Printf("%v: ", car.GetName())
		printCarStatus(car.GetCurrentPosition())
	}
	fmt.Println("-------------------------")
}

func printCarStatus(carPosition int) {
	for i := 0; i < carPosition; i++ {
		fmt.Print(MOVEMENT)
	}
	fmt.Println()
}
