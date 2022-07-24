package api

import (
	"github.com/gin-gonic/gin"
	"uploader/controllers"
	"uploader/ent"
)

func (api *API) register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userSchema := &ent.User{}
		api.bindJsonToStruct(userSchema, ctx)
		statusCode, response := controllers.RegisterUser(api.Crud, userSchema)
		ctx.IndentedJSON(statusCode, response)
	}
}

func (api *API) login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userSchema := &ent.User{}
		api.bindJsonToStruct(userSchema, ctx)
		statusCode, response := controllers.LoginUser(api.Crud, userSchema)
		ctx.IndentedJSON(statusCode, response)
	}
}
