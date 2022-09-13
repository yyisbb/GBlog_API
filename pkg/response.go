package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseJsonOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Code": ERROR_OK,
		"Msg":  GetErrorMsg(ERROR_OK),
	})
}

func ResponseJsonOKAndData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"Code": ERROR_OK,
		"Msg":  GetErrorMsg(ERROR_OK),
		"Data": data,
	})
}

func ResponseJsonError(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"Code": code,
		"Msg":  GetErrorMsg(code),
	})
}
