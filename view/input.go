package view

import "fmt"

var (
	numCars   int
	numRounds int
)

func InputCarNumbers() int {
	fmt.Println("자동차 대수는 몇 대 인가요?")
	fmt.Scanln(&numCars)
	fmt.Println()

	return numCars
}

func InputStageNumbers() int {
	fmt.Println("시도할 회수는 몇 회 인가요?")
	fmt.Scanln(&numRounds)
	fmt.Println()

	return numRounds
}
