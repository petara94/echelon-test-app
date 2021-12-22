package executer

import (
	"log"
)

var MainMachine *Machine

func StartMachine() {
	var err error
	MainMachine, err = AutoStartMachine()

	if err != nil {
		log.Fatal(err)
	}
}
