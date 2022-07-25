package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"uploader/api"
	"uploader/db"
)

func init() {
	os.Clearenv()
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
	}
}

func main() {
	crudInstance, cancel := db.MakeDBReady()
	defer crudInstance.Client.Close()
	defer cancel()
	apiInstance := api.New(&crudInstance, gin.Default())
	apiInstance.RunAPI()
}
