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
	err := router.RunTLS(":"+executer.PORT, "./server.crt", "./server.key")
	if err != nil {
		log.Fatal(err)
	}

}
