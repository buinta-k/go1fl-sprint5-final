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
	datastring = strings.TrimSpace(datastring)

	if strings.Contains(datastring, " ,") || strings.Contains(datastring, ", ") {
		return fmt.Errorf("invalid format")
	}

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
	t.Steps = steps

	if data[1] != "Бег" && data[1] != "Ходьба" {
		return fmt.Errorf("invalid training type")
	}
	t.TrainingType = data[1]

	durStr := data[2]

	if strings.Contains(durStr, " ") {
		return fmt.Errorf("invalid duration")
	}
	if strings.Contains(durStr, ".") {
		return fmt.Errorf("invalid duration")
	}

	dur, err := time.ParseDuration(durStr)
	if err != nil {
		return fmt.Errorf("invalid duration")
	}

	if dur <= 0 {
		return fmt.Errorf("invalid duration")
	}

	t.Duration = dur

	return nil
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
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		dist,
		speed,
		calories,
	)

	return result, nil
}
