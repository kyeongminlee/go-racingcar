package domain

import "fmt"

var MINIMUM_CAR_NUMBER = 0

type Cars []Car

func MakeCars(numCars int) (Cars, error) {
	if numCars <= MINIMUM_CAR_NUMBER {
		return nil, fmt.Errorf("0보다 작은 수입니다.")
	}

	return make([]Car, numCars), nil
}

func (c Cars) Run() Cars {
	for index, car := range c {
		c[index] = car.Run(MakeRandomMovement())
	}

	return c
}
