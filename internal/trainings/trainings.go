package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) error {
	data := strings.Split(datastring, ",")

	if len(data) != 3 {
		return fmt.Errorf("invalid format")
	}

	for i := range data {
		data[i] = strings.TrimSpace(data[i])
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil || steps <= 0 {
		return fmt.Errorf("invalid steps")
	}

	if data[1] == "" {
		return fmt.Errorf("invalid training type")
	}

	str := data[2]

	var hours, minutes int

	n, err := fmt.Sscanf(str, "%dh%dm", &hours, &minutes)
	if err == nil && n == 2 {
		t.Steps = steps
		t.TrainingType = data[1]
		t.Duration = time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute
		return nil
	}

	n, err = fmt.Sscanf(str, "%dh", &hours)
	if err == nil && n == 1 {
		t.Steps = steps
		t.TrainingType = data[1]
		t.Duration = time.Duration(hours) * time.Hour
		return nil
	}

	n, err = fmt.Sscanf(str, "%dm", &minutes)
	if err == nil && n == 1 {
		t.Steps = steps
		t.TrainingType = data[1]
		t.Duration = time.Duration(minutes) * time.Minute
		return nil
	}

	return fmt.Errorf("invalid duration")
}

func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, speed, calories )

	return result, nil
}
