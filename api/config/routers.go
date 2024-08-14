package config

import "github.com/gin-gonic/gin"

type Router interface {
	Register(*gin.RouterGroup)
}
