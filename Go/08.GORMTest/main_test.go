package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestHTTP(t *testing.T) {
	httpResponseString := `{ "status": "expected service response"}`
	req := httptest.NewRequest("GET", "http://localhost:9999/test", nil)
	w := httptest.NewRecorder()
	HandleHello(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("status error, got %v, want %v", status, http.StatusOK)
	}
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != httpResponseString {
		t.Errorf("handle error , got %v, want %v", w.Body.String(), httpResponseString)
	}
}

func TestGetAllUsers(t *testing.T) {
	httpResponse := Response{
		Status: true,
		Data:   []string{"test", "test"},
	}
	req := httptest.NewRequest("GET", "http://localhost:9999/users", nil)
	w := httptest.NewRecorder()
	GetAllUsers(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("status error, got %v, want %v", status, http.StatusOK)
	}
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Info(string(body))
	var res Response
	json.Unmarshal(body, &res)
	if res.Status != httpResponse.Status {
		t.Errorf("handle error , got %v, want %v", w.Body.String(), httpResponse)
	}
}

func TestCreateUser(t *testing.T) {
	httpResponseString := `{"status": "success"}`
	req := httptest.NewRequest("POST", "http://localhost:9999/user/test/test@test.com", nil)
	w := httptest.NewRecorder()
	CreateUser(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("status error, got %v, want %v", status, http.StatusOK)
	}
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != httpResponseString {
		t.Errorf("handle error , got %v, want %v", w.Body.String(), httpResponseString)
	}
}
