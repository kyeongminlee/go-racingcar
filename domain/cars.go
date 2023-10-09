package domain

import (
	"errors"
	"strings"
)

var START_POSITION = 0
var CAR_NAME_LIMIT = 5
var DELIMITER = ","

type Cars []Car

func MakeCars(carNames string) (Cars, error) {
	var newCars []Car
	var err error

	splittedNames := strings.Split(carNames, DELIMITER)

	for _, splittedName := range splittedNames {
		splittedName = strings.TrimSpace(splittedName)
		err = checkNameLength(splittedName)
		if err != nil {
			return nil, err
		}
		newCars = append(newCars, Car{splittedName, START_POSITION})
	}

	return newCars, err
}

func checkNameLength(name string) error {
	if len(name) > CAR_NAME_LIMIT {
		return errors.New("자동차 이름은 5글자 초과할 수 없습니다.")
	}
	return nil
}

func (c Cars) Run(movementStrategy MovementStrategy) Cars {
	for index, car := range c {
		c[index] = car.Run(movementStrategy)
	}

	return c
}

func (c Cars) FindWinner() []string {
	var maxPosition int
	var winnerCars []string

	for _, car := range c {
		if maxPosition < car.position {
			maxPosition = car.position
			winnerCars = []string{car.GetName()}

		} else if maxPosition == car.position {
			winnerCars = append(winnerCars, car.GetName())
		}
	}

	return winnerCars
}
