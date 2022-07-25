package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"uploader/db/crud"
	"uploader/ent"
	"uploader/minio"
)

//
func UploadFile(crud *crud.Crud, username string, file multipart.File, header *multipart.FileHeader) (int, gin.H) {
	user, err := crud.GetUserWithUsername(username)
	if err != nil {
		return http.StatusNotFound, gin.H{
			"message": "user not found",
			"error":   err.Error(),
		}
	}

	url := os.Getenv("URL")

	fileEntity := &ent.FileEntity{
		Type: header.Header["Content-Type"][0],
		Size: header.Size,
		Name: header.Filename,
		URL:  url + fmt.Sprintf("/files/%s", header.Filename),
	}

	createdFile, err := crud.CreateFile(fileEntity, user)
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"message": "can not create file",
			"error":   err.Error(),
		}
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		panic(err)
	}

	bucketName := username

	obj, err := minio.Upload(fileEntity, buf, bucketName)
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"message": "error while uploading file to minio",
			"error":   err.Error(),
		}
	}

	tokens := strings.Split(createdFile.Type, "/")
	fileType := tokens[1]

	return http.StatusCreated, gin.H{
		"id":       createdFile.ID,
		"file_url": fileEntity.URL,
		"type":     fileType,
		"size":     obj.Size,
	}
}

func GetAllFiles(crud *crud.Crud, username string) (int, gin.H) {
	bucketName := username
	allFiles, err := crud.GetAllFiles(username)
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "error while fetching file entities from database",
		}
	}

	return http.StatusOK, gin.H{
		"message": fmt.Sprintf("all files in bucket with name %s", bucketName),
		"files":   allFiles,
	}
}

func DeleteFileWithID(crud *crud.Crud, fileID int, username string) (int, gin.H) {
	bucketName := username
	deletedFileEntity, err := crud.DeleteFileWithID(username, fileID)
	if deletedFileEntity == nil {
		return http.StatusNotFound, gin.H{
			"message": "no file entity found with given id",
		}
	}
	if err != nil {
		if strings.Contains(err.Error(), "user") {
			return http.StatusNotFound, gin.H{
				"message": "user not found",
				"error":   err.Error(),
			}
		} else if strings.Contains(err.Error(), "file") {
			return http.StatusInternalServerError, gin.H{
				"message": "error while deleting file from database",
				"error":   err.Error(),
			}
		}
	}
	err = minio.DeleteFileWithName(bucketName, deletedFileEntity.Name)
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"message": "error while deleting object from minio bucket",
			"error":   err.Error(),
		}
	}
	return http.StatusOK, gin.H{
		"message": "file deleted successfully",
		"file":    deletedFileEntity,
	}
}

func Download() {

}
