package executer

import (
	"bytes"
	"errors"
	"os/exec"
	"runtime"
)

// RequestBody Структура для приходящих пакетов на выполнение
type RequestBody struct {
	CMD   string `json:"cmd"`
	OS    string `json:"os"`
	Stdin string `json:"stdin"`
}

type ExecResult struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}
type BadExecResult struct {
	Request *RequestBody `json:"command"`
	Error   string       `json:"error"`
}

type Machine struct {
	OS string `json:"os"`
}

func (m Machine) Exec(cmd, stdin string) (*ExecResult, error) {

	runner := exec.Command("sh")
	runner.Stdin = bytes.NewReader([]byte(cmd + "\n" + stdin + "\nexit"))

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	runner.Stdout = &stdout
	runner.Stderr = &stderr

	err := runner.Run()

	if err != nil {
		return nil, err
	}

	return &ExecResult{stdout.String(), stderr.String()}, nil
}

func NewMachine(OS string) *Machine {
	return &Machine{OS}
}

func StartLinuxMachine() *Machine {
	return NewMachine(LINUX_OS)
}

func StartWinMachine() *Machine {
	return NewMachine(WINDOWS_OS)
}

func AutoStartMachine() (*Machine, error) {
	if runtime.GOOS == WINDOWS_OS {
		return StartWinMachine(), nil
	} else if runtime.GOOS == LINUX_OS {
		return StartLinuxMachine(), nil
	} else {
		return NewMachine(ERROR_OS), errors.New(ERROR_INIT_OS)
	}
}

func NewBadExecResult(request *RequestBody, error string) *BadExecResult {
	return &BadExecResult{Request: request, Error: error}
}
