package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":        "mall",
		"using":       "gin",
		"description": "Mall is a online shopping platform with [gin](https://github.com/gin-gonic/gin).",
	})
}
