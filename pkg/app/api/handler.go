package api

import (
	"awesomeProject1/pkg/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func RouteNotFound(c *gin.Context) {
	SendResponse(c, errno.RouterNotFound, nil)
	return
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(fmt.Sprintf("----------------panic: %v, %s\n", r, string(debug.Stack())))
			SendResponse(c, errno.InternalServerError, nil)
			return
		}
	}()
	c.Next()
}
