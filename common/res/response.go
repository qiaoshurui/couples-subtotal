// description:
// @author renshiwei
// Date: 2022/10/5 21:30

package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code: SUCCESS_CODE,
		Msg:  SUCCESS_MSG,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: SUCCESS_CODE,
		Msg:  SUCCESS_MSG,
		Data: data,
	})
}
func OkWithMessage(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: SUCCESS_CODE,
		Msg:  msg,
		Data: data,
	})
}
func ErrorDefault(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code: ERROR_CODE,
		Msg:  ERROR_MSG,
	})
}
func ParamError(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code: ERROR_CODE,
		Msg:  PARAM_ERROR,
	})
}
func Error(c *gin.Context, code, msg string) {
	c.JSON(http.StatusOK, &Response{
		Code: ERROR_CODE,
		Msg:  msg,
	})
}

func ErrorData(c *gin.Context, code, msg string, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: ERROR_CODE,
		Msg:  msg,
		Data: data,
	})
}
