package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function

func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
	totalFodder, err := fc.FodderAmount(numCows)
	if err != nil {
		return 0, err
	}

	ff, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	fodderPerCow := totalFodder * ff / float64(numCows)

	return fodderPerCow, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function

func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, numCows)
}

// TODO: define the 'ValidateNumberOfCows' function

type InvalidCowsError struct {
	message string
	numCows int
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}

func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{
			message: "there are no negative cows",
			numCows: numCows,
		}
	} else if numCows == 0 {
		return &InvalidCowsError{
			message: "no cows don't need food",
			numCows: numCows,
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
