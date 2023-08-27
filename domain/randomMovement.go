package domain

import "math/rand"

type MovementStrategy interface {
	Movement() bool
}

var MAX_VALUE = 10
var CONDITION = 4

type RandomMovement struct {
}

func MakeRandomMovement() *RandomMovement {
	return &RandomMovement{}
}

func (r RandomMovement) Movement() bool {
	return rand.Intn(MAX_VALUE) >= CONDITION
}
