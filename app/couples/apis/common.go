package apis

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func getPageInfo(c *gin.Context) (int, int) {
	//获取分页参数
	pageStr := c.Query("offset")
	sizeStr := c.Query("limit")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return int(page), int(size)
}
