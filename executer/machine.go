package executer

import (
	"log"
)

var MainMachine *Machine

// StartMachine инициализация исполнителя команд
func StartMachine() {
	var err error
	MainMachine, err = AutoStartMachine()

	if err != nil {
		log.Fatal(err)
	}
}
