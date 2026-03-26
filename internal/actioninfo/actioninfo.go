package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, action := range dataset {
		if err := dp.Parse(action); err != nil {
			log.Printf("не получилось получить информацию о тренировке: %v", err)
			continue
		}
		if str, err := dp.ActionInfo(); err != nil {
			log.Printf("не получилось получить информацию о тренировке: %v", err)
		} else {
			fmt.Println(str)
		}
	}
}
