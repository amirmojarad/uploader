package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
	"uploader/api"
	"uploader/db/crud"
	"uploader/ent"
	"uploader/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
)

type TestClient struct {
	context.Context
	ent.Client
	testing.T
	context.CancelFunc
	api.API
}

func GetFiles(tc *TestClient, uploadUrl string, token string) *httptest.ResponseRecorder {
	method := "GET"
	response := httptest.NewRecorder()
	request, err := http.NewRequest(method, uploadUrl, nil)
	if err != nil {
		tc.Error(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	tc.Engine.ServeHTTP(response, request)
	return response

}

type UploadFileResponse struct {
	FileUrl string `json:"file_url"`
	Id      int    `json:"id"`
	Size    int    `json:"size"`
	Type    string `json:"type"`
}

func UploadFile(tc *TestClient, uploadUrl string, token string, ufr *UploadFileResponse) *httptest.ResponseRecorder {
	method := "POST"

	payloadFile := &bytes.Buffer{}
	writer := multipart.NewWriter(payloadFile)
	file, errFile1 := os.Open("5750684.jpg")
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("file", filepath.Base("5750684.jpg"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return nil
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, uploadUrl, payloadFile)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	tc.Engine.ServeHTTP(w, req)

	wBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if err := json.Unmarshal(wBytes, &ufr); err != nil {
		panic(err)
	}
	return w
}

type RegisterSchema struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	User    struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"user"`
}

func CreateUser(tc *TestClient, registerUrl string, obj *RegisterSchema) {
	response := httptest.NewRecorder()

	testUser := ent.User{Username: "testusername", Password: "testpass123"}
	payload, _ := json.Marshal(testUser)
	request, _ := http.NewRequest("POST", registerUrl, bytes.NewBuffer(payload))
	tc.Engine.ServeHTTP(response, request)
	if err := json.Unmarshal(response.Body.Bytes(), &obj); err != nil {
		panic(err)
	}
}

func (testClient TestClient) CallCancelAndClose() {
	testClient.CancelFunc()
	testClient.Client.Close()
}

func GetTestClientAndContext(t *testing.T) *TestClient {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*24)

	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %+v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema changes: %+v", err)
	}

	crudInstance := crud.Crud{
		Client: client,
		Ctx:    &ctx,
	}

	apiInstance := api.New(&crudInstance, gin.Default())
	apiInstance.SetUpAPI()
	return &TestClient{Context: ctx, Client: *client, T: *t, CancelFunc: cancel, API: apiInstance}
}
