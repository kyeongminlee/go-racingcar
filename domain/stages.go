package domain

import "fmt"

var MINIMUM_STAGE_NUMBER = 0
var UNIT_STAGE = 1

type Stages struct {
	stage int
}

func MakeStage(numStage int) (*Stages, error) {
	if numStage <= MINIMUM_STAGE_NUMBER {
		return nil, fmt.Errorf("0보다 작은 수입니다.")
	}

	return &Stages{
		stage: numStage,
	}, nil
}

func (s *Stages) StartGame(cars Cars, movementStrategy MovementStrategy) Cars {
	s.countDown()
	cars = cars.Run(movementStrategy)

	return cars
}

func (s *Stages) IsFinished() bool {
	return s.stage == MINIMUM_STAGE_NUMBER
}

func (s *Stages) CurrentStage() int {
	return s.stage
}

func (s *Stages) countDown() {
	s.stage -= UNIT_STAGE
}
