package main

import (
	"fmt"
	"math/rand"
	"racingcar/domain"
	"racingcar/view"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// numCars := view.InputCarNumbers()
	carNames := view.InputCarNames()
	numStage := view.InputStageNumbers()

	cars, err := domain.MakeCars(carNames)

	stages, err := domain.MakeStage(numStage)
	if err != nil {
		fmt.Println(err)
	}

	view.PrintResultMention()
	for {
		if stages.IsFinished() {
			break
		}
		cars := stages.StartGame(cars, domain.MakeRandomMovement())
		view.PrintCarsStatus(cars)
	}

}
