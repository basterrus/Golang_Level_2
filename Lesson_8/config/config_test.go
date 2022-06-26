package config

import (
	"log"
	"testing"
)

func TestInit(t *testing.T) {
	cfg, err := Init()
	if cfg == nil {
		log.Fatalf("Ошибка: файл конфигурации не загружен")
	}

	if err != nil {
		log.Fatalf("Ошибка: файл конфигурации не загружен: %s", err)
	}
}
