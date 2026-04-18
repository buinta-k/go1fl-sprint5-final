package actioninfo


import (
	"fmt" // Добавь эту строку
)

type DataParser interface {
	Parse(datastring string) error
    ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	    for _, data := range dataset {
        
        err := dp.Parse(data)
        if err != nil {
            fmt.Printf("ошибка обработки данных '%s': %v\n", data, err)
            continue
        }

        
        info, err := dp.ActionInfo()
        if err != nil {
            fmt.Printf("ошибка обработки данных '%s': %v\n", data, err)
            continue
        }

        fmt.Println(info)
    }
}
