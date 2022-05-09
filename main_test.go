package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignI(t *testing.T) {
	var body = `{"username":"user1","password":"password1"}`

	req := httptest.NewRequest("POST", "http://localhost:8000/signin", strings.NewReader(body))

	rec := httptest.NewRecorder()

	Signin(rec, req)

	if rec.Code != 200 {
		t.Errorf("Expected status code of 200, got %d", rec.Code)
	}

	data, _ := ioutil.ReadAll(rec.Body)
	if len(data) == 0 || data == nil {
		t.Error("An error occured while generating access token")
	}
}
