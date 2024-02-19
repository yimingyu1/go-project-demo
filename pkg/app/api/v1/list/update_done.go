package list

import (
	"awesomeProject1/internal/list/service"
	"awesomeProject1/pkg/app/api"
	"github.com/gin-gonic/gin"
	"log"
)

func UpdateDone(c *gin.Context) {
	var updateDoneRequest UpdateTodoRequest
	if err := c.ShouldBindJSON(&updateDoneRequest); err != nil {
		if err := c.ShouldBindJSON(&updateDoneRequest); err != nil {
			log.Printf("updateTodoDone bind params err: %v", err)
			api.SendResponse(c, err, nil)
			return
		}
	}
	service.Svc.ListSvc.UpdateDone(c, updateDoneRequest.Id)
	api.SendResponse(c, nil, "更新待办事务完成")
}
