package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
)

type Minio struct {
	Ctx             context.Context
	MinioClient     *minio.Client
	Endpoint        string
	accessKeyID     string
	secretAccessKey string
}

func NewMinio() *Minio {
	minioInstance := &Minio{
		Ctx:             context.Background(),
		Endpoint:        os.Getenv("MINIO_ENDPOINT"),
		accessKeyID:     os.Getenv("MINIO_ACCESSKEY"),
		secretAccessKey: os.Getenv("MINIO_SECRETKEY"),
	}
	minioClient, errInit := minio.New(minioInstance.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioInstance.accessKeyID, minioInstance.secretAccessKey, ""),
		Secure: false,
	})
	if errInit != nil {
		log.Fatalln(errInit)
	}
	minioInstance.MinioClient = minioClient
	return minioInstance
}

func (minioInstance Minio) ConnectToBucket(bucketName string) {
	// )
	location := "us-east-1"

	err := minioInstance.MinioClient.MakeBucket(minioInstance.Ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioInstance.MinioClient.BucketExists(minioInstance.Ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}
