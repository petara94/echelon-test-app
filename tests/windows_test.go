package tests

import (
	"bytes"
	"echelon-test-app/executer"
	"echelon-test-app/route"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWinHttp1(t *testing.T) {
	executer.StartMachine()
	r := route.InitRouter()

	w := httptest.NewRecorder()

	Rb := executer.RequestBody{
		CMD:   "ping 0.0.0.0 -c 1",
		OS:    "windows",
		Stdin: "",
	}

	bodyBytes, err := json.Marshal(Rb)
	body := bytes.NewReader(bodyBytes)

	assert.Nil(t, err)
	req, err := http.NewRequest("GET", "/api/v1/exec", body)

	assert.Nil(t, err)
	r.ServeHTTP(w, req)

	// Вывод на разных машинах отличатся сильно,
	// но если программа была запущена, то код будет 200
	assert.Equal(t, w.Code, 200)
}

func TestWinHttp2(t *testing.T) {
	executer.StartMachine()
	r := route.InitRouter()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/api/v1/os", nil)

	assert.Nil(t, err)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Body.String(), "{\"os\":\"windows\"}")
}

func TestWinHttp3(t *testing.T) {
	executer.StartMachine()
	r := route.InitRouter()

	w := httptest.NewRecorder()

	Rb := executer.RequestBody{
		CMD:   "..\\stdin_test_prog\\test_stdin.exe",
		OS:    "windows",
		Stdin: "2 2\n",
	}

	bodyBytes, err := json.Marshal(Rb)
	body := bytes.NewReader(bodyBytes)

	assert.Nil(t, err)
	req, err := http.NewRequest("GET", "/api/v1/exec", body)

	assert.Nil(t, err)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Body.String(), "{\"stdout\":\"a + a * b = 6\\r\\n\",\"stderr\":\"\"}")
}
