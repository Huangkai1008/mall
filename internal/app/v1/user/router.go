package user

import (
	"github.com/gin-gonic/gin"

	"mall/internal/pkg/router"
)

func NewRouter(h *Handler) router.Group {
	return func(r *gin.RouterGroup) {
		userApi := r.Group("/users")
		{
			userApi.POST("register", h.Register)
		}
	}
}
