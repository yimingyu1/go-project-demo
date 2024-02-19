package list

import (
	"awesomeProject1/internal/list/service"
	"awesomeProject1/pkg/app/api"
	"github.com/gin-gonic/gin"
	"log"
)

func Delete(c *gin.Context) {
	var deleteTodoRequest DeleteTodoRequest
	if err := c.ShouldBindJSON(&deleteTodoRequest); err != nil {
		log.Printf("deleteTodo bind params err: %v", err)
		api.SendResponse(c, err, nil)
		return
	}
	service.Svc.ListSvc.DeletedTodo(c, deleteTodoRequest.Id)
	api.SendResponse(c, nil, "删除成功")
}
