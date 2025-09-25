package utils

import (
	"flint/security"

	"github.com/gin-gonic/gin"
)

type Route interface {
	Route() (Method, Path, *security.Policy)
	Do(c *gin.Context)
}

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
)

type Path string
