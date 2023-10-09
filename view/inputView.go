package view

import "fmt"

var (
	carNames  string
	numRounds int
)

// func InputCarNumbers() int {
// 	fmt.Println("자동차 대수는 몇 대 인가요?")
// 	fmt.Scanln(&numCars)
// 	fmt.Println()

// 	return numCars
// }

func InputCarNames() string {
	fmt.Println("경주할 자동차 이름을 입력하세요(이름은 쉼표(,)를 기준으로 구분).")
	fmt.Scanln(&carNames)
	fmt.Println()

	return carNames
}

func InputStageNumbers() int {
	fmt.Println("시도할 회수는 몇 회 인가요?")
	fmt.Scanln(&numRounds)
	fmt.Println()

	return numRounds
}
