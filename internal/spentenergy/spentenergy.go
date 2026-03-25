package spentenergy

import (
	"time"

	"github.com/Yandex-Practicum/tracker/internal/errs"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := errs.ValidateStep(steps); err != nil {
		return 0, err
	}
	if err := errs.ValidateDuration(duration); err != nil {
		return 0, err
	}
	if weight <= 0 {
		return 0, errs.ErrInvalidWeight
	}
	if height <= 0 {
		return 0, errs.ErrInvalidHeight
	}
	speedKmh := MeanSpeed(steps, height, duration)
	calories := (walkingCaloriesCoefficient * weight * speedKmh * duration.Minutes()) / minInH
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := errs.ValidateStep(steps); err != nil {
		return 0, err
	}
	if err := errs.ValidateDuration(duration); err != nil {
		return 0, err
	}
	if weight <= 0 {
		return 0, errs.ErrInvalidWeight
	}
	if height <= 0 {
		return 0, errs.ErrInvalidHeight
	}
	speedKmh := MeanSpeed(steps, height, duration)
	calories := weight * speedKmh * duration.Minutes() / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	speedKmh := dist / duration.Hours()
	return speedKmh
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
