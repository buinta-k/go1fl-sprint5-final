package trainings

import (
		"github.com/buinta-k/go1fl-sprint5-final/tree/main/internal/personaldata",
		"strconv"
	   )

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	dataString:=strings.Split(datastring,",")
	if dataString!=3 {
		return fmt.Errorf("Слайс не равен трем")
	}
	Steps,err =strconv.Atoi(dataString[0]) 
		if err!=nil {
		return 0, err
	}
	TrainingType = dataString[1]
	Duration, _:=time.ParseDuration(dataString[3])
	
}

func (t Training) ActionInfo() (string, error) {
	
func (t Training) ActionInfo() (string, error) {
    dist := Distance(t.Steps, t.Height)
    speed := MeanSpeed(t.Steps, t.Height, t.Duration)
    
    var calories float64
    var err error

    switch t.TrainingType {
    case "Бег":
        calories, err = RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
    case "Ходьба":
        calories, err = WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
    default:
        return "", fmt.Errorf("неизвестный тип тренировки")
    }

    if err != nil {
        return "", err
    }

    result := fmt.Sprintf(
        "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
        t.TrainingType,
        t.Duration.Hours(),
        dist,
        speed,
        calories,
    )

    return result, nil
}

}
