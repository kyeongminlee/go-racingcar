package view

import (
	"fmt"
	"racingcar/domain"
	"strings"
)

var MOVEMENT = "-"
var DELIMITER = ", "

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

func PrintWinners(carNames []string) {
	fmt.Printf("%v가 최종 우승했습니다.", strings.Join(carNames, DELIMITER))
}
