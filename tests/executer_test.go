package tests

import (
	"executer_server/executer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMachine_exec(t *testing.T) {

	mach, err := executer.AutoStartMachine()

	assert.Nil(t, err)

	res, err := mach.Exec("echo 123 | wc", "")

	assert.Nil(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, res.Stdout, "      1       1       4\n")
}

func TestMachine_exec2(t *testing.T) {

	mach, err := executer.AutoStartMachine()

	assert.Nil(t, err)

	res, err := mach.Exec("../stdini_test_prog/test_stdin", "2\n2\n")

	assert.Nil(t, err)

	assert.Equal(t, res.Stdout, "a + a * b = 6\n")
}

func TestMachine_execFail(t *testing.T) {

	mach, err := executer.AutoStartMachine()

	if err != nil {
		t.Error(err.Error())
	}

	_, err = mach.Exec("lss", "")

	assert.NotNil(t, err)
}
