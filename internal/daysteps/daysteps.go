package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	data := strings.Split(datastring, ",")

	if len(data) != 2 {
		return fmt.Errorf("invalid format")
	}

	for i := range data {
		data[i] = strings.TrimSpace(data[i])
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil || steps <= 0 {
		return fmt.Errorf("invalid steps")
	}

	str := data[1]

	var hours, minutes int

	n, err := fmt.Sscanf(str, "%dh%dm", &hours, &minutes)
	if err == nil && n == 2 {
		ds.Steps = steps
		ds.Duration = time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute
		return nil
	}

	n, err = fmt.Sscanf(str, "%dh", &hours)
	if err == nil && n == 1 {
		ds.Steps = steps
		ds.Duration = time.Duration(hours) * time.Hour
		return nil
	}

	n, err = fmt.Sscanf(str, "%dm", &minutes)
	if err == nil && n == 1 {
		ds.Steps = steps
		ds.Duration = time.Duration(minutes) * time.Minute
		return nil
	}

	return fmt.Errorf("invalid duration")
}

func (ds *DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
