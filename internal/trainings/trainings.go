package trainings

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/errs"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return errs.ErrInvalidFormat
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	if err = errs.ValidateStep(steps); err != nil {
		return err
	}
	action := parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}

	if err = errs.ValidateDuration(duration); err != nil {
		return err
	}
	t.Steps = steps
	t.TrainingType = action
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	default:
		return "", errs.ErrUnknownType
	}
	speedKmh := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	dist := spentenergy.Distance(t.Steps, t.Height)
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		dist,
		speedKmh,
		calories,
	), nil
}
