package main

import (
	"echelon-test-app/executer"
	"echelon-test-app/route"
	"log"
)

func main() {

	// Инициализация сервера и исполнителя команд
	executer.StartMachine()
	router := route.InitRouter()

	// Запуск сервера
	err := router.RunTLS(":"+executer.PORT, "./ssl/server.crt", "./ssl/server.key")
	if err != nil {
		log.Fatal(err)
	}
}
