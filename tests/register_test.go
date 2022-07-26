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

const RegisterUrl = "/api/register"

func TestRegisterBadRequest(t *testing.T) {
	tc := GetTestClientAndContext(t)
	defer tc.CallCancelAndClose()

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", RegisterUrl, nil)
	tc.Engine.ServeHTTP(w, request)
	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestRegisterDuplicate(t *testing.T) {
	tc := GetTestClientAndContext(t)
	defer tc.CallCancelAndClose()

	w := httptest.NewRecorder()

	testUser := ent.User{Username: "testUsername", Password: "testpass123"}
	payload, _ := json.Marshal(testUser)
	request, _ := http.NewRequest("POST", RegisterUrl, bytes.NewBuffer(payload))
	tc.Engine.ServeHTTP(w, request)
	assert.Equal(t, http.StatusCreated, w.Code)

	w = httptest.NewRecorder()

	testUser = ent.User{Username: "testUsername", Password: "testpass123"}
	payload, _ = json.Marshal(testUser)
	request, _ = http.NewRequest("POST", RegisterUrl, bytes.NewBuffer(payload))
	tc.Engine.ServeHTTP(w, request)
	assert.Equal(t, http.StatusConflict, w.Code)
	log.Println(w.Code)
	log.Println(w.Body)
}

func TestRegister(t *testing.T) {
	tc := GetTestClientAndContext(t)
	defer tc.CallCancelAndClose()

	w := httptest.NewRecorder()

	testUser := ent.User{Username: "testUsername", Password: "testpass123"}
	payload, _ := json.Marshal(testUser)
	request, _ := http.NewRequest("POST", RegisterUrl, bytes.NewBuffer(payload))
	tc.Engine.ServeHTTP(w, request)
	assert.Equal(t, http.StatusCreated, w.Code)
	rs := RegisterSchema{}
	if err := json.Unmarshal(w.Body.Bytes(), &rs); err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, rs.Message)
	assert.NotEmpty(t, rs.User)
	assert.NotEmpty(t, rs.Token)
}
