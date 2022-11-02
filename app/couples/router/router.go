package router

import (
	"github.com/gin-gonic/gin"
)

func InitCoupleRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api/v1")
	RegisterUserRouter(v1)
	RelationshipRouter(v1)
	PhotoRouter(v1)
	PhotoAlbumRouter(v1)
	DynamicRouter(v1)
	return r
}
