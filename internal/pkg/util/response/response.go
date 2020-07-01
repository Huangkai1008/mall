package response

import (
	"github.com/gin-gonic/gin"
	"mall/internal/pkg/ecode"
	"net/http"
)

// Created return with `Created` status code and created entity schema json.
func Created(c *gin.Context, schema interface{}) {
	c.JSON(http.StatusCreated, schema)
}

// BadEntityRequest abort with `UnprocessableEntity` status code and error json.
func BadEntityRequest(c *gin.Context, errs ecode.MallError) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, errs)
}

// BadRequest abort with `BadRequest` status code and error message json.
func BadRequest(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}
