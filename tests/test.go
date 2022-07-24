package tests

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"os"
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
