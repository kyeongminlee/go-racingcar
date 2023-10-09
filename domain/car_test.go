package domain

import (
	"testing"
)

type MockMovementStrategy struct {
	shouldMove bool
}

func (m MockMovementStrategy) Movement() bool {
	return m.shouldMove
}

func TestMakeCar(t *testing.T) {
	tests := []struct {
		description        string
		car                Car
		movementStrategy   MovementStrategy
		expectedPosition   int
		expectedShouldMove bool
	}{
		{
			description:        "Car should move",
			car:                Car{name: "Car1", position: 0},
			movementStrategy:   MockMovementStrategy{shouldMove: true},
			expectedPosition:   1,
			expectedShouldMove: true,
		},
		{
			description:        "Car should not move",
			car:                Car{name: "Car2", position: 0},
			movementStrategy:   MockMovementStrategy{shouldMove: false},
			expectedPosition:   0,
			expectedShouldMove: false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := test.car.Run(test.movementStrategy)

			if result.position != test.expectedPosition {
				t.Errorf("Expected position %d, but got %d", test.expectedPosition, result.position)
			}
		})
	}
}
