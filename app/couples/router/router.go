package router

import (
	"github.com/gin-gonic/gin"
)

func InitCoupleRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api/v1")
	RegisterUserRouter(v1)
	RelationshipRouter(v1)
	return r
}
