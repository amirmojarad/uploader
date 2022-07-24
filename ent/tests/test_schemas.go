package tests

import (
	"context"
	"entgo.io/ent/examples/fs/ent"
	"entgo.io/ent/examples/fs/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"testing"
	"uploader/db"
)

func setUp(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("cannot make enttest client: %+v", err)
		}
	}(client)
	ctx, cancel := db.GetDatabaseContext()
	defer (*cancel)()
	if err := client.Schema.WriteTo(*ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %+v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed craeting schemas: %+v", err)
	}
	client = (*client).Debug()

}

func TestUser(t *testing.T) {

}
