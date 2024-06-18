package httphandler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"syscall"
	"testing"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var res interface{}
	err = json.Unmarshal(rawBody, &res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(res)

	w.Write([]byte("Hello, wooooooorld!"))
}

var url = "http://test.ru"

func TestHandler(t *testing.T) {
	test := `
		hello": "world"
		"a": ["b",1,["c",2],{"d":3}]]
	}`
	rawReq, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	syscall.Setenv("METHOD", "POST")
	aa := bytes.NewBuffer(rawReq)
	req := httptest.NewRequest("__METHOD__", url, aa)
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	Handler(w, req)
	rawBytes, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rawBytes))
}
