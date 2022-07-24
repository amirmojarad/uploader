package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"uploader/db/crud"
	"uploader/ent"
	"uploader/usecases"
)

func LoginUser(crud *crud.Crud, userSchema *ent.User) (int, gin.H) {
	if fetchedUser, err := crud.GetUserWithUsername(userSchema.Username); err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("user with username %s not found.", userSchema.Username),
				"error":   err.Error(),
			}
		}
		return http.StatusInternalServerError, gin.H{
			"message": "error while fetching user from database",
			"error":   err.Error(),
		}
	} else {
		log.Println("ASDSADSADSAD")
		log.Println(usecases.CheckPasswordHash(userSchema.Password, fetchedUser.Password))
		if !usecases.CheckPasswordHash(userSchema.Password, fetchedUser.Password) {
			return http.StatusBadRequest, gin.H{
				"message": "wrong password",
			}
		} else {
			return http.StatusOK, gin.H{
				"user":    fetchedUser,
				"message": "logged in successfully",
			}
		}
	}
}

func RegisterUser(crud *crud.Crud, userSchema *ent.User) (int, gin.H) {
	if createdUser, err := crud.CreateUser(userSchema); err != nil {
		if strings.Contains(err.Error(), "username") {
			return http.StatusConflict, gin.H{
				"message": "error while adding user to database",
				"error":   err.Error(),
			}
		}
		return http.StatusInternalServerError, gin.H{
			"message": "error while adding user to database",
			"error":   err.Error(),
		}
	} else {
		return http.StatusCreated, gin.H{
			"user":    createdUser,
			"message": "user created successfully",
		}
	}
}
