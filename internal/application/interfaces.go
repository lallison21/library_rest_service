package application

import "github.com/gin-gonic/gin"

type StatusHandler interface {
	Ping() gin.HandlerFunc
}

type AuthHandler interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
}
