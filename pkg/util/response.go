package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
	})
}

func OkData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": data,
	})
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "fail",
	})
}

func FailMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
	})
}

func CodeMsg(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
