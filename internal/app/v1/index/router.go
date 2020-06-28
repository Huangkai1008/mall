package index

import "github.com/gin-gonic/gin"

func NewRouter(h *Handler) func(r *gin.RouterGroup) {
	return func(r *gin.RouterGroup) {
		indexAPI := r.Group("/")
		{
			indexAPI.GET("ping", Ping)
		}
	}
}
