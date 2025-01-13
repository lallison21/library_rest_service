package application

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
}
