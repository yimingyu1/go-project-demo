package list

import (
	"awesomeProject1/internal/list/service"
	"awesomeProject1/pkg/app/api"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	list := service.Svc.ListSvc.GetList(c)
	api.SendResponse(c, nil, list)
}
