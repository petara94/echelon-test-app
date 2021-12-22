package tests

import (
	"bytes"
	"echelon-test-app/executer"
	"echelon-test-app/route"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLinuxExec1(t *testing.T) {

	mach, err := executer.AutoStartMachine()

	assert.Nil(t, err)

	res, err := mach.Exec("ping 0.0.0.0 -c 1", "")

	assert.Nil(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, strings.Split(res.Stdout, "\n")[0], "PING 0.0.0.0 (127.0.0.1) 56(84) bytes of data.")
}

func TestMachineExec2(t *testing.T) {

	mach, err := executer.AutoStartMachine()

	assert.Nil(t, err)

	res, err := mach.Exec("echo 123 | wc", "")

	assert.Nil(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, res.Stdout, "      1       1       4\n")
}

func TestHttp1(t *testing.T) {
	executer.StartMachine()
	r := route.InitRouter()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/api/v1/os", nil)

	assert.Nil(t, err)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Body.String(), "{\"os\":\"linux\"}")
}

func TestHttp2(t *testing.T) {
	executer.StartMachine()
	r := route.InitRouter()

	w := httptest.NewRecorder()

	Rb := executer.RequestBody{
		CMD:   "echo 123",
		OS:    "linux",
		Stdin: "",
	}

	bodyBytes, err := json.Marshal(Rb)
	body := bytes.NewReader(bodyBytes)

	assert.Nil(t, err)
	req, err := http.NewRequest("GET", "/api/v1/exec", body)

	assert.Nil(t, err)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Body.String(), "{\"stdout\":\"123\\n\",\"stderr\":\"\"}")
}

func TestHttp3(t *testing.T) {
	executer.StartMachine()
	r := route.InitRouter()

	w := httptest.NewRecorder()

	Rb := executer.RequestBody{
		CMD:   "echo 123 | wc",
		OS:    "linux",
		Stdin: "",
	}

	bodyBytes, err := json.Marshal(Rb)
	body := bytes.NewReader(bodyBytes)

	assert.Nil(t, err)
	req, err := http.NewRequest("GET", "/api/v1/exec", body)

	assert.Nil(t, err)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Body.String(), "{\"stdout\":\"      1       1       4\\n\",\"stderr\":\"\"}")
}

func TestHttp4(t *testing.T) {
	executer.StartMachine()
	r := route.InitRouter()

	w := httptest.NewRecorder()

	Rb := executer.RequestBody{
		CMD:   "",
		OS:    "windows",
		Stdin: "",
	}

	bodyBytes, err := json.Marshal(Rb)
	body := bytes.NewReader(bodyBytes)

	assert.Nil(t, err)
	req, err := http.NewRequest("GET", "/api/v1/exec", body)

	assert.Nil(t, err)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 400)
}
