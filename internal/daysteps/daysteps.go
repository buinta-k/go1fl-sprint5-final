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

func (ds *DaySteps) Parse(datastring string) (err error) {
    dataString := strings.Split(datastring, ",")
    if len(dataString) != 3 {
        return fmt.Errorf("invalid format")
    }

    steps, err := strconv.Atoi(dataString[0])
    if err != nil || steps <= 0 { 
        return fmt.Errorf("invalid steps")
    }

    duration, err := time.ParseDuration(dataString[2])
    if err != nil || duration <= 0 { 
        return fmt.Errorf("invalid duration")
    }

    t.Steps = steps
    t.TrainingType = dataString[1]
    t.Duration = duration
    return nil
	
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps, distance, calories)

	return result, nil
}
