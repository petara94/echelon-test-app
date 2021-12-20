package executer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMachine_exec(t *testing.T) {

	mach, err := AutoStartMachine()

	if err != nil {
		t.Error(err.Error())
	}

	res, err := mach.Exec("echo 123 | wc", "")

	assert.Nil(t, err)

	assert.Equal(t, res.Stdout, "      1       1       4\n")
}
