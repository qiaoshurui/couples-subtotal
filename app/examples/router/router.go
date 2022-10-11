// description:
// @author renshiwei
// Date: 2022/10/5 21:58

package router

import (
	"github.com/gin-gonic/gin"
)

func InitHelloRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api/v1")
	RegisterHelloRouter(v1)
	return r
}
