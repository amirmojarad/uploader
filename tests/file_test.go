package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"uploader/minio"
)

const UploadURL = "/api/upload"

func TestUpload(t *testing.T) {
	loadEnv()

	tc := GetTestClientAndContext(t)
	defer tc.CallCancelAndClose()
	rs := RegisterSchema{}

	CreateUser(tc, RegisterUrl, &rs)

	token := rs.Token
	ufr := UploadFileResponse{}

	response := UploadFile(tc, UploadURL, token, &ufr)

	fileEntity, err := tc.Crud.GetFile(rs.User.Username, ufr.Id)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, ufr.Id, fileEntity.ID)

	minioObj, err := minio.GetFileFromBucket(rs.User.Username, fileEntity.Name)
	if err != nil {
		t.Error(err)
	}
	objInfo, err := minioObj.Stat()

	assert.Equal(t, fileEntity.Name, objInfo.Key)
}

type GetAllFilesResponse struct {
	Files []struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Type  string `json:"type"`
		Size  int    `json:"size"`
		Url   string `json:"url"`
		Edges struct {
		} `json:"edges"`
	} `json:"files"`
	Message string `json:"message"`
}

func TestGet(t *testing.T) {
	loadEnv()

	tc := GetTestClientAndContext(t)
	defer tc.CallCancelAndClose()

	rs := &RegisterSchema{}
	CreateUser(tc, RegisterUrl, rs)
	token := rs.Token

	ufr := UploadFileResponse{}
	UploadFile(tc, UploadURL, token, &ufr)

	response := GetFiles(tc, UploadURL, token)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
		return
	}
	t.Log(string(body))

	gafr := GetAllFilesResponse{}

	if err := json.Unmarshal(body, &gafr); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, 1, len(gafr.Files))
}
