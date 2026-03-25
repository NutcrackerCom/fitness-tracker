package errs

import (
	"errors"
	"time"
)

var (
	ErrInvalidWeight    = errors.New("неправильный вес")
	ErrInvalidHeight    = errors.New("неправильный рост")
	ErrInvalidFormat    = errors.New("неправильный формат")
	ErrNoSteps          = errors.New("нулевое количество шагов")
	ErrNegativeSteps    = errors.New("отрицательное число шагов")
	ErrNegativeDuration = errors.New("отрицательная продолжительность")
	ErrZeroDuration     = errors.New("нулевая продолжительность")
	ErrUnknownType      = errors.New("неизвестный тип тренировки")
)

func ValidateStep(steps int) error {
	if steps == 0 {
		return ErrNoSteps
	}
	if steps < 0 {
		return ErrNegativeSteps
	}
	return nil
}

func ValidateDuration(duration time.Duration) error {
	if duration < 0 {
		return ErrNegativeDuration
	}
	if duration == 0 {
		return ErrZeroDuration
	}
	return nil
}
