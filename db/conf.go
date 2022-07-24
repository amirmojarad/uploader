package db

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
	"uploader/db/crud"
	"uploader/ent"
)

type databaseConf struct {
	User     string
	DbName   string
	Host     string
	Password string
}

func createDatabaseConf() *databaseConf {
	return &databaseConf{
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PASSWORD"),
	}
}

func MakeDBReady() (crud.Crud, context.CancelFunc) {
	dbConf := createDatabaseConf()
	client, err := ent.Open("postgres",
		fmt.Sprintf("user=%s dbname=%s host=%s password=%s sslmode=disable",
			dbConf.User,
			dbConf.DbName,
			dbConf.Host,
			dbConf.Password,
		),
	)
	if err != nil {
		log.Printf("error while opening database connection: %+v", err)
	}

	client = (*client).Debug()

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*24)

	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %+v", err)
	}
	// create ent schemas into database if there are not exists in db. this function uses context.Background.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema changes: %+v", err)
	}
	return crud.Crud{
		Ctx:    &ctx,
		Client: client,
	}, cancel
}
