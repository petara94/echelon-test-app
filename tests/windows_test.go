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
