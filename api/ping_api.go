package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/etcd_backend/internal"
)

func init() {
	internal.AddRoute(internal.Routes{
		{
			Method:  http.MethodGet,
			Path:    "/etcd/ping",
			Handler: ping,
		},
	})
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
