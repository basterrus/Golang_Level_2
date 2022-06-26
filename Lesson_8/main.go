package main

import (
	"fmt"
	"github.com/basterrus/Golang_Level_2/Lesson_8/config"
	"github.com/basterrus/Golang_Level_2/Lesson_8/internal"
	"log"
)

func main() {

	// Инициализируем конфигурацию приложения
	conf, err := config.Init()
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла конфигурации: %s", err)
	}

	// Инициализируем подсистему флагов
	err = conf.InitFlags()
	if err != nil {
		log.Fatalf("Ошибка при инициализации флагов: %s", err)
	}
	// Запускаем функцию сканирования каталогов
	err = internal.SearchFiles(conf)
	if err != nil {
		fmt.Println(err)
	}

}
