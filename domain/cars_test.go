package domain

import (
	"errors"
	"reflect"
	"testing"
)

func TestMakeCars(t *testing.T) {
	tests := []struct {
		input       string
		expected    Cars
		expectedErr error
		description string
	}{
		{
			input: "Car1,Car2,Car3",
			expected: Cars{
				Car{name: "Car1", position: START_POSITION},
				Car{name: "Car2", position: START_POSITION},
				Car{name: "Car3", position: START_POSITION},
			},
			expectedErr: nil,
			description: "Multiple car names",
		},
		// 추가 테스트 케이스를 여기에 추가할 수 있습니다.
		{
			input: "Car1, Car2, Car3",
			expected: Cars{
				Car{name: "Car1", position: START_POSITION},
				Car{name: "Car2", position: START_POSITION},
				Car{name: "Car3", position: START_POSITION},
			},
			expectedErr: nil,
			description: "Multiple car names",
		},
		{
			input: "Cara1, Carb2, Carc3",
			expected: Cars{
				Car{name: "Cara1", position: START_POSITION},
				Car{name: "Carb2", position: START_POSITION},
				Car{name: "Carc3", position: START_POSITION},
			},
			expectedErr: nil,
			description: "Multiple car names",
		},
		{
			input:       "Cara1,Carab2",
			expected:    nil,
			expectedErr: errors.New("자동차 이름은 5글자 초과할 수 없습니다."),
			description: "Multiple car names",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result, err := MakeCars(test.input)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}

			if (err != nil && test.expectedErr == nil) || (err == nil && test.expectedErr != nil) {
				t.Errorf("Expected error \"%v\", but got \"%v\"", test.expectedErr, err)
			}
		})
	}
}

func TestFindWinner(t *testing.T) {
	tests := []struct {
		description string
		cars        Cars
		expected    []string
	}{
		{
			description: "Single winner",
			cars: Cars{
				Car{name: "Car1", position: 2},
				Car{name: "Car2", position: 3},
				Car{name: "Car3", position: 1},
			},
			expected: []string{"Car2"},
		},
		{
			description: "Multiple winners",
			cars: Cars{
				Car{name: "Car1", position: 2},
				Car{name: "Car2", position: 3},
				Car{name: "Car3", position: 3},
				Car{name: "Car4", position: 1},
			},
			expected: []string{"Car2", "Car3"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := test.cars.FindWinner()

			// 우승자 슬라이스의 길이 검증
			if len(result) != len(test.expected) {
				t.Errorf("Expected %d winner(s), but got %d", len(test.expected), len(result))
			}

			// 우승자의 이름과 위치 검증
			for i, name := range result {
				if name != test.expected[i] {
					t.Errorf("Expected winner %v, but got %v", test.expected[i], name)
				}
			}
		})
	}
}
