package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"uploader/ent"
)

const LoginUrl = "/api/login"

func TestLoginBadRequest(t *testing.T) {
	tc := GetTestClientAndContext(t)
	defer tc.CallCancelAndClose()

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", LoginUrl, nil)
	tc.Engine.ServeHTTP(w, request)
	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestLoginNotFound(t *testing.T) {
	tc := GetTestClientAndContext(t)
	defer tc.CallCancelAndClose()

	w := httptest.NewRecorder()

	testUser := ent.User{Username: "testUsername", Password: "testpass123"}
	payload, _ := json.Marshal(testUser)
	request, _ := http.NewRequest("POST", RegisterUrl, bytes.NewBuffer(payload))
	tc.Engine.ServeHTTP(w, request)
	assert.Equal(t, http.StatusCreated, w.Code)

	w = httptest.NewRecorder()

	testUser = ent.User{Username: "testUsasdername", Password: "testpass123"}
	payload, _ = json.Marshal(testUser)
	request, _ = http.NewRequest("POST", LoginUrl, bytes.NewBuffer(payload))
	tc.Engine.ServeHTTP(w, request)
	assert.Equal(t, http.StatusNotFound, w.Code)
	log.Println(w.Body)
}

func TestLoginWrongPassword(t *testing.T) {
	tc := GetTestClientAndContext(t)
	defer tc.CallCancelAndClose()

	w := httptest.NewRecorder()

	testUser := ent.User{Username: "testUsername", Password: "testpass123"}
	payload, _ := json.Marshal(testUser)
	request, _ := http.NewRequest("POST", RegisterUrl, bytes.NewBuffer(payload))
	tc.Engine.ServeHTTP(w, request)
	assert.Equal(t, http.StatusCreated, w.Code)

	w = httptest.NewRecorder()

	testUser = ent.User{Username: "testUsername", Password: "wrongpassword123"}
	payload, _ = json.Marshal(testUser)
	request, _ = http.NewRequest("POST", LoginUrl, bytes.NewBuffer(payload))
	tc.Engine.ServeHTTP(w, request)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogin(t *testing.T) {
	tc := GetTestClientAndContext(t)
	defer tc.CallCancelAndClose()

	w := httptest.NewRecorder()

	testUser := ent.User{Username: "testUsername", Password: "testpass123"}
	payload, _ := json.Marshal(testUser)
	request, _ := http.NewRequest("POST", RegisterUrl, bytes.NewBuffer(payload))
	tc.Engine.ServeHTTP(w, request)
	assert.Equal(t, http.StatusCreated, w.Code)

	w = httptest.NewRecorder()

	payload, _ = json.Marshal(testUser)
	request, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
	tc.Engine.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
}
