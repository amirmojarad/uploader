package minio

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"uploader/ent"
)

func Upload(fileSchema *ent.FileEntity, buffer *bytes.Buffer, bucketName string) (*minio.UploadInfo, error) {
	minioInstance := NewMinio()
	minioInstance.ConnectToBucket(bucketName)
	object, err := minioInstance.MinioClient.PutObject(
		context.Background(),
		bucketName,
		fileSchema.Name,
		buffer,
		fileSchema.Size,
		minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return nil, err
	}
	return &object, nil
}

func GetAllFilesFromBucket(bucketName string) ([]*minio.ObjectInfo, error) {
	minioInstance := NewMinio()
	minioInstance.ConnectToBucket(bucketName)
	opts := minio.ListObjectsOptions{
		UseV1:     true,
		Recursive: true,
	}
	objects := make([]*minio.ObjectInfo, 1)
	for object := range minioInstance.MinioClient.ListObjects(context.Background(), bucketName, opts) {
		if object.Err != nil {
			return nil, object.Err
		}
		objects = append(objects, &object)
	}
	return objects, nil
}

func DeleteFileWithName(bucketName, fileName string) error {
	minioInstance := NewMinio()
	minioInstance.ConnectToBucket(bucketName)
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err := minioInstance.MinioClient.RemoveObject(context.Background(), bucketName, fileName, opts)
	if err != nil {
		return err
	}
	return nil
}
