package main

import (
	"echelon-test-app/executer"
	"log"
)

var machine *executer.Machine

func StartMachine() {
	var err error
	machine, err = executer.AutoStartMachine()

	if err != nil {
		log.Fatal(err)
	}
}
