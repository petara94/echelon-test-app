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
	err := router.RunTLS(":"+executer.PORT, "./tls/server.crt", "./tls/server.key")
	if err != nil {
		log.Fatal(err)
	}
}
