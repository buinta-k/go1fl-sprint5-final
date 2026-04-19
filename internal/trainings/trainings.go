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

func (t *DaySteps) Parse(datastring string) error {
	data := strings.Split(datastring, ",")
	if len(data) != 2 { // В daysteps обычно 2 параметра
		return fmt.Errorf("invalid format")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil || steps <= 0 {
		return fmt.Errorf("invalid steps")
	}

	durStr := strings.TrimSpace(data[1])
	dur, err := time.ParseDuration(durStr)
	if err != nil || dur <= 0 {
		return fmt.Errorf("invalid duration")
	}

	t.Steps = steps
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
