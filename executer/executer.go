package executer

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
	"runtime"
	"strings"
)

// RequestBody Структура для приходящих пакетов на выполнение
type RequestBody struct {
	CMD   string `json:"cmd"`
	OS    string `json:"os"`
	Stdin string `json:"stdin"`
}

// ExecResult Структура для удачного выполнения
type ExecResult struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

// BadExecResult Структура для вывода ошибок выполнения
type BadExecResult struct {
	Request *RequestBody `json:"command"`
	Error   string       `json:"error"`
}

type Machine struct {
	OS string `json:"os"`
}

func (m Machine) Exec(cmd, stdin string) (*ExecResult, error) {

	commands := strings.Split(cmd, "|")

	var stderr bytes.Buffer

	var runner *exec.Cmd

	// Выполнение с перенаправлением stdout в stdin
	for _, command := range commands {
		// Удаление пробелов по краям
		command = strings.Trim(command, " ")

		runner = exec.Command(strings.Split(command, " ")[0],
			strings.Split(command, " ")[1:]...)

		runner.Stdin = bytes.NewReader([]byte(stdin))
		runner.Stderr = &stderr

		outToIn, err := runner.StdoutPipe()

		if err != nil {
			return nil, err
		}

		err = runner.Start()

		if err != nil {
			return nil, err
		}

		stdinByte, _ := io.ReadAll(outToIn)
		stdin = string(stdinByte)

		_ = runner.Wait()
	}

	return &ExecResult{stdin, stderr.String()}, nil
}

// NewMachine Конструктор Machine
func NewMachine(OS string) *Machine {
	return &Machine{OS}
}

// StartLinuxMachine Конструктор linux Machine
func StartLinuxMachine() *Machine {
	return NewMachine(LINUX_OS)
}

// StartWinMachine Конструктор windows Machine
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
