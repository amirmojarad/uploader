package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"uploader/controllers"
)

func (api API) upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := api.getUsername(ctx)
		file, header, err := ctx.Request.FormFile("file")
		defer file.Close()
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "invalid file type or data",
			})
			log.Println(err.Error())
			return
		}
		statusCode, jsonResponse := controllers.UploadFile(api.Crud, username, file, header)
		ctx.IndentedJSON(statusCode, jsonResponse)

	}
}

func (api API) getAllFiles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := api.getUsername(ctx)
		statusCode, jsonResponse := controllers.GetAllFiles(api.Crud, username)
		ctx.IndentedJSON(statusCode, jsonResponse)
	}
}

func (api API) delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := api.getUsername(ctx)
		fileID, err := strconv.Atoi(ctx.Param("file_id"))
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "url does not contain any file id",
			})
			return
		}
		statusCode, jsonResponse := controllers.DeleteFileWithID(api.Crud, fileID, username)
		ctx.IndentedJSON(statusCode, jsonResponse)
	}
}
