package tests

import (
	"bytes"
	"echelon-test-app/executer"
	"echelon-test-app/route"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWinExec1(t *testing.T) {

	mach, err := executer.AutoStartMachine()

	assert.Nil(t, err)

	_, err = mach.Exec("ping 0.0.0.0 -c 1", "")

	assert.Nil(t, err)
	if err != nil {
		return
	}

	//assert.Equal(t, strings.Split(res.Stdout, "\n")[0], "PING 0.0.0.0 (127.0.0.1) 56(84) bytes of data.")
}

func TestWinExec2(t *testing.T) {

	mach, err := executer.AutoStartMachine()

	assert.Nil(t, err)

	_, err = mach.Exec("ping 0.0.0.0 -c 1", "")

	assert.Nil(t, err)
	if err != nil {
		return
	}

	//assert.Equal(t, strings.Split(res.Stdout, "\n")[0], "PING 0.0.0.0 (127.0.0.1) 56(84) bytes of data.")
}

func TestWinHttp1(t *testing.T) {
	executer.StartMachine()
	r := route.InitRouter()

	w := httptest.NewRecorder()

	Rb := executer.RequestBody{
		CMD:   "ping 0.0.0.0 -c 1",
		OS:    "linux",
		Stdin: "",
	}

	bodyBytes, err := json.Marshal(Rb)
	body := bytes.NewReader(bodyBytes)

	assert.Nil(t, err)
	req, err := http.NewRequest("GET", "/api/v1/exec", body)

	assert.Nil(t, err)
	r.ServeHTTP(w, req)

	log.Println(w.Body.String())

	//assert.Equal(t, w.Body.String(), "")
}
