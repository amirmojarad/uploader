package tests

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"uploader/minio"
)

func loadEnv() {
	os.Clearenv()
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
	}
}

func TestMinioConnection(t *testing.T) {

}

func TestBucketExists(t *testing.T) {
	loadEnv()

	minioInstance := minio.NewMinio()
	bucketName := os.Getenv("MINIO_BUCKET")
	minioInstance.ConnectToBucket(bucketName)

	ok, err := minioInstance.MinioClient.BucketExists(minioInstance.Ctx, bucketName)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, true, ok)
}
