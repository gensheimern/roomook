package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)



type LoggedInUser struct {
	User string `json:"user"`
}



func  GetLoggedInUser(ctx *gin.Context) {

		var loggedInUser LoggedInUser
		loggedInUser.User = ctx.Request.Header.Get("x-authorized-user")
		ctx.JSON(http.StatusOK, loggedInUser)
}
