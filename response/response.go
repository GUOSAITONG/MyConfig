package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"操作成功"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, msg string, Data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 200,
		"msg":  msg,
		"data": Data,
	})
}
func Error(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
