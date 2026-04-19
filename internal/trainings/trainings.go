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
func (t *Training) Parse(datastring string) error {
	data := strings.Split(datastring, ",")
    
	if len(data) != 3 {
		return fmt.Errorf("invalid format")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil || steps <= 0 {
		return fmt.Errorf("invalid steps")
	}

	trainingType := strings.TrimSpace(data[1])
	if trainingType == "" {
		return fmt.Errorf("invalid training type")
	}

	durStr := strings.TrimSpace(data[2])
	dur, err := time.ParseDuration(durStr)
	if err != nil || dur <= 0 {
		return fmt.Errorf("invalid duration")
	}

	t.Steps = steps
	t.TrainingType = trainingType
	t.Duration = dur

	return nil
}
