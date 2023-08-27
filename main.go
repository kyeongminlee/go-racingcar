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

	numCars := view.InputCarNumbers()
	numStage := view.InputStageNumbers()

	cars, err := domain.MakeCars(numCars)
	if err != nil {
		fmt.Println(err)
	}

	stages, err := domain.MakeStage(numStage)
	if err != nil {
		fmt.Println(err)
	}

	view.PrintResultMention()
	for {
		if stages.IsFinished() {
			break
		}
		cars := stages.StartGame(cars)
		view.PrintCarsStatus(cars)
	}

}
