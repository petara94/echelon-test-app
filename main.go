package main

import (
	"echelon-test-app/executer"
	"log"
)

func main() {

	// Инициализация сервера и исполнителя команд
	StartMachine()
	router := InitRouter()

	// Запуск сервера
	err := router.Run(":" + executer.PORT)
	if err != nil {
		log.Fatal(err)
	}

}
