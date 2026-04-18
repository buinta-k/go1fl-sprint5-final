package daysteps

import ("strings"
		"strconv"
	   "github.com/buinta-k/go1fl-sprint5-final/blob/main/internal/spentenergy/spentenergy.go"
	   )

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal 
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataString := strings.Split(datastring, ",")
    if len(dataString) < 2 {
        return fmt.Errorf("Некорректный формат")
    }

	ds.Steps, err = strconv.Atoi(dataString[0])
	if err != nil {
		return fmt.Errorf("Ошибка обработки данных: %w", err)
	}

	ds.Duration, err = time.ParseDuration(dataString[1])
    if err != nil {
        return fmt.Errorf("Ошибка обработки данных: %w", err)
    }
    return nil
	
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	
    if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\n Дистанция: %d км.\n Вы сожгли %d ккал.\n", steps, distanсe, calories)

	return result, nil
}
