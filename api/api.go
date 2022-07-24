/*
api package contains endpoints
*/

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"uploader/api/auth"
	"uploader/db/crud"
)

func New(crudInstance *crud.Crud, ginEngine *gin.Engine) API {
	return API{
		Crud:       crudInstance,
		Engine:     ginEngine,
		jwtService: auth.JWTAuthService(),
	}
}

type API struct {
	Crud       *crud.Crud
	Engine     *gin.Engine
	jwtService auth.JWTService
}

func (api API) bindJsonToStruct(givenStruct interface{}, ctx *gin.Context) interface{} {
	if err := ctx.BindJSON(&givenStruct); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "given json schema is not valid",
			"err":     err.Error(),
		})
	}
	return givenStruct
}

func (api API) SetUpAPI() {
	api.Engine.POST("/api/register", api.register())
	api.Engine.POST("/api/login", api.login())
}

func (api API) RunAPI() {
	api.SetUpAPI()
	api.Engine.Run(":8080")
}
