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

func (d *DaySteps) Parse(datastring string) error {
    data := strings.Split(datastring, ",")
    if len(data) != 2 { // В daysteps обычно 2 параметра: шаги и время
        return fmt.Errorf("Некорректный формат")
    }

    steps, err := strconv.Atoi(data[0]) 
    if err != nil || steps <= 0 {
        return fmt.Errorf("Некорректный формат")
    }

    durStr := strings.TrimSpace(data[1])
    // УБЕРИТЕ проверку на точку здесь!
    duration, err := time.ParseDuration(durStr)
    if err != nil || duration <= 0 {
        return fmt.Errorf("Некорректный формат")
    }

    d.Steps = steps
    d.Duration = duration
    return nil
}

	return fmt.Errorf("Некорректный формат")
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
